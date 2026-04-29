package v1

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"spark-oj/api/v1/problem"
	"spark-oj/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func (c *ControllerProblem) GetTestcases(ctx context.Context, req *problem.GetTestcasesReq) (res *problem.GetTestcasesRes, err error) {
	res = &problem.GetTestcasesRes{}

	cnt, err := dao.Problem.Ctx(ctx).Count("problem_id", req.ProblemId)
	if err != nil || cnt == 0 {
		g.Log().Error(ctx, err, cnt)
		return nil, err
	}

	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, req.ProblemId)
	if !gfile.Exists(problemDir) {
		return res, nil
	}

	inputFiles, err := gfile.ScanDirFile(problemDir, "*.in", false)
	if err != nil {
		return nil, err
	}
	sort.Strings(inputFiles)

	for _, inputFile := range inputFiles {
		baseName := strings.TrimSuffix(filepath.Base(inputFile), ".in")
		outputFile := filepath.Join(problemDir, baseName+".out")
		if !gfile.Exists(outputFile) {
			continue
		}
		res.Testcases = append(res.Testcases, problem.TestcaseItem{
			Name:       baseName,
			InputSize:  gfile.Size(inputFile),
			OutputSize: gfile.Size(outputFile),
		})
	}

	return res, nil
}
