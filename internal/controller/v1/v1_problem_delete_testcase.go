package v1

import (
	"context"
	"path/filepath"
	"spark-oj/api/v1/problem"
	"spark-oj/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func (c *ControllerProblem) DeleteTestcase(ctx context.Context, req *problem.DeleteTestcaseReq) (res *problem.DeleteTestcaseRes, err error) {
	res = &problem.DeleteTestcaseRes{}

	cnt, err := dao.Problem.Ctx(ctx).Count("problem_id", req.ProblemId)
	if err != nil || cnt == 0 {
		g.Log().Error(ctx, err, cnt)
		return nil, err
	}

	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, req.ProblemId)

	name := filepath.Base(req.Name)
	inFile := filepath.Join(problemDir, name+".in")
	outFile := filepath.Join(problemDir, name+".out")

	if gfile.Exists(inFile) {
		if err = gfile.Remove(inFile); err != nil {
			return nil, err
		}
	}
	if gfile.Exists(outFile) {
		if err = gfile.Remove(outFile); err != nil {
			return nil, err
		}
	}

	return res, nil
}
