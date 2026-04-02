package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) GetInfo(ctx context.Context, req *problem.GetInfoReq) (res *problem.GetInfoRes, err error) {
	res = &problem.GetInfoRes{}
	md := dao.Problem.Ctx(ctx)

	d := &entity.Problem{}
	err = md.Where("problemId", req.ProblemId).Scan(d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(d, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
