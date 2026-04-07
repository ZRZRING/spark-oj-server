package v1

import (
	"context"
	"path/filepath"
	"spark-oj/internal/dao"

	"spark-oj/api/v1/core"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerCore) UploadTestCase(ctx context.Context, req *core.UploadTestCaseReq) (res *core.UploadTestCaseRes, err error) {
	res = &core.UploadTestCaseRes{}

	if req.TestCases == nil {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "TestCases is nil")
		g.Log().Error(ctx, err)
		return nil, err
	}

	md := dao.Problem.Ctx(ctx)
	cnt, err := md.Count("problem_id", req.ProblemId)
	if err != nil || cnt == 0 {
		g.Log().Error(ctx, err, cnt)
		return nil, err
	}

	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, gconv.String(req.ProblemId))
	if err = gfile.Mkdir(problemDir); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	for _, file := range req.TestCases {
		filename, err := file.Save(problemDir)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		g.Log().Debug(ctx, "Saved file:", filename)
	}

	res.Path = problemDir
	return res, nil
}
