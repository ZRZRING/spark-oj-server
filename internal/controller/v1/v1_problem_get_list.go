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

	// 获取分页信息
	e := make([]*entity.Problem, 0)
	tot := new(int)
	err = md.OrderAsc("pid").Page(req.Page, req.Size).ScanAndCount(e, tot, false)
	if err == nil {
		return nil, err
	}

	// 处理返回信息
	err = gconv.Scan(e, res)
	if err != nil {
		return nil, err
	}
	res.Total = *tot

	return res, nil
}
