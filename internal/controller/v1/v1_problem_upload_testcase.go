package v1

import (
	"context"
	"path/filepath"
	"spark-oj/api/v1/problem"
	"spark-oj/internal/dao"
	"spark-oj/pkg/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerProblem) UploadTestcase(ctx context.Context, req *problem.UploadTestcaseReq) (res *problem.UploadTestcaseRes, err error) {
	res = &problem.UploadTestcaseRes{}

	if req.Testcases == nil {
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

	// 检查文件
	ioMap := make(map[string]int)
	for _, file := range req.Testcases {
		if file.Size > consts.MAX_FILE_SIZE {
			err = gerror.NewCode(gcode.CodeInvalidRequest, "File size is too large")
			g.Log().Error(ctx, err)
			return nil, err
		}
		// 扩展名为 in 和 out，且文件名相同的文件必须成对出现
		ext := filepath.Ext(file.Filename)
		if ext == ".in" {
			ioMap[filepath.Base(file.Filename)] |= 1
		} else if ext == ".out" {
			ioMap[filepath.Base(file.Filename)] |= 2
		} else {
			err = gerror.NewCode(gcode.CodeInvalidRequest, "File extension is invalid")
			g.Log().Error(ctx, err)
			return nil, err
		}
	}
	for filename, cnt := range ioMap {
		if cnt == 1 {
			err = gerror.NewCodef(gcode.CodeInvalidRequest, "File %s.out doesn't exist", filepath.Base(filename))
			g.Log().Error(ctx, err)
			return nil, err
		} else if cnt == 2 {
			err = gerror.NewCodef(gcode.CodeInvalidRequest, "File %s.in doesn't exist", filepath.Base(filename))
			g.Log().Error(ctx, err)
			return nil, err
		}
	}

	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, gconv.String(req.ProblemId))
	if err = gfile.Mkdir(problemDir); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	for _, file := range req.Testcases {
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
