package v1

import (
	"context"
	"fmt"

	"spark-oj/api/v1/problem"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"
	"spark-oj/pkg/enums"
)

func (c *ControllerProblem) GetPage(ctx context.Context, req *problem.GetPageReq) (res *problem.GetPageRes, err error) {
	var (
		e   []*entity.Problem
		tot int
	)

	err = dao.Problem.Ctx(ctx).
		OrderAsc("problem_id").
		Page(req.Page, req.Size).
		ScanAndCount(&e, &tot, false)
	if err != nil {
		return nil, err
	}

	res = &problem.GetPageRes{Total: tot}
	for _, p := range e {
		submitCount, _ := dao.Submission.Ctx(ctx).
			Where("problem_id", p.ProblemId).
			Distinct().
			Fields("username").
			Count()
		acceptCount, _ := dao.Submission.Ctx(ctx).
			Where("problem_id", p.ProblemId).
			Where("result", "Accepted").
			Distinct().
			Fields("username").
			Count()

		res.Problems = append(res.Problems, &problem.GetPageItem{
			ProblemId:   fmt.Sprintf("%d", p.ProblemId),
			Title:       p.Title,
			JudgeType:   enums.JudgeType(p.JudgeType),
			Rating:      p.Rating,
			AcceptCount: acceptCount,
			SubmitCount: submitCount,
			CreateBy:    p.CreateBy,
			CreateTime:  p.CreateAt.Layout("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}
