package problem

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spoj/api/problem/v1"
	"spoj/internal/dao"
	"spoj/internal/model/entity"
)

func (c *ControllerV1) GetProblem(ctx context.Context, req *v1.GetProblemReq) (res *v1.GetProblemRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := dao.Problem.Ctx(ctx)
	problem := &entity.Problem{}
	pid := gconv.String(r.Get("pid").Val())
	err = md.Where("pid", pid).Scan(&problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	err = gconv.Scan(problem, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return res, nil
}
