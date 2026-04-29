package v1

import (
	"context"
	"fmt"

	"spark-oj/api/v1/problem"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type problemCountRow struct {
	ProblemId   int `json:"problem_id"`
	SubmitCount int `json:"submit_count"`
	AcceptCount int `json:"accept_count"`
}

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

	ids := make([]int, 0, len(e))
	for _, p := range e {
		ids = append(ids, p.ProblemId)
	}

	var counts []problemCountRow
	err = g.DB().Ctx(ctx).Raw(`
		SELECT problem_id,
			COUNT(DISTINCT CASE WHEN result = 'Accepted' THEN username END) AS accept_count,
			COUNT(DISTINCT username) AS submit_count
		FROM submission
		WHERE problem_id IN(?) AND delete_at IS NULL
		GROUP BY problem_id
	`, ids).Scan(&counts)
	if err != nil {
		return nil, err
	}

	countMap := make(map[int]*problemCountRow, len(counts))
	for i := range counts {
		countMap[counts[i].ProblemId] = &counts[i]
	}

	res = &problem.GetPageRes{Total: tot}
	for _, p := range e {
		var acceptCount, submitCount int
		if c, ok := countMap[p.ProblemId]; ok {
			acceptCount = c.AcceptCount
			submitCount = c.SubmitCount
		}
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
