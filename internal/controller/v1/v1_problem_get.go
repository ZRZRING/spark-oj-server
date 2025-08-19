package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) Get(ctx context.Context, req *problem.GetReq) (res *problem.GetRes, err error) {
	res = &problem.GetRes{}
	md := dao.Problem.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	pid := gconv.String(r.Get("pid").Val())

	d := &entity.Problem{}
	err = md.Where("pid", pid).Scan(d)
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
