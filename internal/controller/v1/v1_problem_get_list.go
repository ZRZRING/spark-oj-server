package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) GetList(ctx context.Context, req *problem.GetListReq) (res *problem.GetListRes, err error) {
	res = &problem.GetListRes{}
	md := dao.Problem.Ctx(ctx)

	e := make([]*entity.Problem, 0)
	tot := 0
	err = md.OrderAsc("pid").Page(req.Page, req.Size).ScanAndCount(&e, &tot, false)
	if err == nil {
		return nil, err
	}

	err = gconv.Scan(e, &res.Problems)
	if err != nil {
		return nil, err
	}
	res.Total = tot

	return res, nil
}
