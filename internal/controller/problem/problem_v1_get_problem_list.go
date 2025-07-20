package problem

import (
	"context"
	"spark-oj-server/internal/consts"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/problem/v1"
)

func (c *ControllerV1) GetProblemList(ctx context.Context, req *v1.GetProblemListReq) (res *v1.GetProblemListRes, err error) {
	res = &v1.GetProblemListRes{}
	md := dao.Problem.Ctx(ctx)
	tot, err := md.Count()
	if err != nil {
		return nil, err
	}
	num := (req.Page - 1) * req.Size
	if tot <= num {
		return nil, consts.ErrInvalidPageSize
	}
	var problems []*entity.Problem
	err = md.
		Fields("pid", "type", "title").
		OrderAsc("pid").
		Limit(req.Size).Offset((req.Page - 1) * req.Size).
		Scan(&problems)
	if err != nil {
		return nil, err
	}
	res.Problems = make([]*v1.Problem, len(problems))
	for i, p := range problems {
		res.Problems[i] = &v1.Problem{
			Pid:   p.Pid,
			Title: p.Title,
		}
	}
	res.Total = tot
	return
}
