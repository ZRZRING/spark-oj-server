package v1

import (
	"context"
	"spark-oj/api/v1/core"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"
	"spark-oj/internal/service"
	"spark-oj/pkg/enums"
	"spark-oj/pkg/utils"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerCore) Judge(ctx context.Context, req *core.JudgeReq) (res *core.JudgeRes, err error) {
	res = &core.JudgeRes{}

	if req.Language == "" {
		req.Language = "cpp"
	}
	if _, ok := service.LanguageConfigs[req.Language]; !ok {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "不支持的语言")
		g.Log().Error(ctx, err, req.Language)
		return nil, err
	}

	problem := &entity.Problem{}
	err = dao.Problem.Ctx(ctx).Where("problem_id", req.ProblemId).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if problem.ProblemId == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "题目不存在")
		g.Log().Error(ctx, err, req.ProblemId)
		return nil, err
	}

	insertData := g.Map{
		"problem_id": req.ProblemId,
		"username":   req.Username,
		"result":     string(enums.JudgeStatusWaiting),
		"language":   req.Language,
		"code":       req.Code,
	}
	if req.ContestId != "" {
		insertData["contest_id"] = req.ContestId
	}

	id, err := dao.Submission.Ctx(ctx).InsertAndGetId(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	submissionId := gconv.String(id)
	res.SubmissionId = submissionId
	res.Result = enums.JudgeStatusWaiting

	go runJudge(gctx.New(), submissionId, req.ProblemId, req.Code, req.Language)

	return res, nil
}

func runJudge(ctx context.Context, submissionId string, problemIdStr string, code string, language enums.Language) {
	problemId := gconv.Int(problemIdStr)

	_, err := dao.Submission.Ctx(ctx).Where("submission_id", submissionId).Data("result", string(enums.JudgeStatusRunning)).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	testCases, err := service.CollectTestCases(problemId)
	if err != nil {
		g.Log().Error(ctx, err)
		dao.Submission.Ctx(ctx).Where("submission_id", submissionId).Data(g.Map{
			"result":      string(enums.JudgeStatusRuntimeError),
			"memory_cost": 0,
			"time_cost":   0,
		}).Update()
		return
	}

	judgeResult := enums.JudgeStatusAccepted
	var maxTime, maxMemory int64

	for _, testCase := range testCases {
		input := gfile.GetContents(testCase.InputPath)
		expected := gfile.GetContents(testCase.OutputPath)

		exeRes, executeErr := service.ExecuteCode(ctx, &service.ExecuteCodeRequest{
			Code:     code,
			Input:    input,
			Language: language,
		})
		if executeErr != nil {
			g.Log().Error(ctx, executeErr)
			judgeResult = enums.JudgeStatusRuntimeError
			break
		}

		if exeRes.Time > maxTime {
			maxTime = exeRes.Time
		}
		if exeRes.Memory > maxMemory {
			maxMemory = exeRes.Memory
		}

		if exeRes.Status != enums.JudgeStatusAccepted {
			judgeResult = exeRes.Status
			break
		}

		if !utils.IsOutputEqual(exeRes.Output, expected) {
			judgeResult = "Wrong Answer"
			break
		}
	}

	_, err = dao.Submission.Ctx(ctx).Where("submission_id", submissionId).Data(g.Map{
		"result":      string(judgeResult),
		"memory_cost": maxMemory,
		"time_cost":   maxTime / (1000 * 1000),
	}).Update()
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
