package v1

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"spark-oj-server/api/v1/core"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/internal/service"

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
	err = dao.Problem.Ctx(ctx).Where("pid", req.Pid).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if problem.Pid == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "题目不存在")
		g.Log().Error(ctx, err, req.Pid)
		return nil, err
	}

	pid := gconv.Int(req.Pid)
	testCases, err := collectTestCases(pid)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	result := "Accepted"
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

		if exeRes.Status != "Accepted" {
			result = exeRes.Status
			break
		}

		if !isOutputEqual(exeRes.Output, expected) {
			result = "Wrong Answer"
			break
		}
	}

	sid, err := dao.Submission.Ctx(ctx).Data(&do.Submission{
		Title:      problem.Title,
		Pid:        req.Pid,
		Username:   req.Username,
		Cid:        req.Cid,
		Result:     result,
		Language:   req.Language,
		MemoryCost: maxMemory,
		TimeCost:   maxTime / (1000 * 1000),
		Code:       req.Code,
	}).InsertAndGetId()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	res.Sid = gconv.String(sid)
	res.Result = result
	return res, nil
}

type judgeTestCase struct {
	InputPath  string
	OutputPath string
}

func collectTestCases(pid int) ([]judgeTestCase, error) {
	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, gconv.String(pid))
	if !gfile.Exists(problemDir) {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "测试用例不存在")
	}

	inputFiles, err := gfile.ScanDirFile(problemDir, "*.in", false)
	if err != nil {
		return nil, err
	}
	if len(inputFiles) == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "未找到输入测试用例")
	}
	sort.Strings(inputFiles)

	testCases := make([]judgeTestCase, 0, len(inputFiles))
	for _, inputFile := range inputFiles {
		outputFile := strings.TrimSuffix(inputFile, ".in") + ".out"
		if !gfile.Exists(outputFile) {
			return nil, gerror.NewCode(gcode.CodeInvalidRequest, "测试用例缺少输出文件")
		}

		testCases = append(testCases, judgeTestCase{
			InputPath:  inputFile,
			OutputPath: outputFile,
		})
	}

	return testCases, nil
}

func isOutputEqual(actual, expected string) bool {
	return strings.TrimSpace(actual) == strings.TrimSpace(expected)
}
