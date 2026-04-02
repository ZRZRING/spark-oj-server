package v1

import (
	"context"
	"spark-oj-server/api/v1/core"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/internal/service"
	"spark-oj-server/pkg/enums"
	"spark-oj-server/pkg/utils"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	err = dao.Problem.Ctx(ctx).Where("problemId", req.ProblemId).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if problem.ProblemId == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "题目不存在")
		g.Log().Error(ctx, err, req.ProblemId)
		return nil, err
	}

	problemId := gconv.Int(req.ProblemId)
	testCases, err := service.CollectTestCases(problemId)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	result := enums.JudgeStatusAccepted
	var maxTime, maxMemory int64

	for _, testCase := range testCases {
		input := gfile.GetContents(testCase.InputPath)
		expected := gfile.GetContents(testCase.OutputPath)

		exeRes, executeErr := service.ExecuteCode(ctx, &service.ExecuteCodeRequest{
			Code:     req.Code,
			Input:    input,
			Language: req.Language,
		})
		if executeErr != nil {
			g.Log().Error(ctx, executeErr)
			return nil, executeErr
		}

		if exeRes.Time > maxTime {
			maxTime = exeRes.Time
		}
		if exeRes.Memory > maxMemory {
			maxMemory = exeRes.Memory
		}

		if exeRes.Status != enums.JudgeStatusAccepted {
			result = exeRes.Status
			break
		}

		if !utils.IsOutputEqual(exeRes.Output, expected) {
			result = "Wrong Answer"
			break
		}
	}

	data := do.Submission{
		ProblemId:  req.ProblemId,
		Username:   req.Username,
		Result:     result,
		Language:   req.Language,
		MemoryCost: maxMemory,
		TimeCost:   maxTime / (1000 * 1000),
		Code:       req.Code,
	}
	if req.ContestId != "" {
		data.ContestId = req.ContestId
	}

	submissionId, err := dao.Submission.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	res.SubmissionId = gconv.String(submissionId)
	res.Result = result
	return res, nil
}
