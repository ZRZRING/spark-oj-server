package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) GetProblem(ctx context.Context, req *problem.GetProblemReq) (res *problem.GetProblemRes, err error) {
	res = &problem.GetProblemRes{}
	md := dao.Problem.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	pid := gconv.String(r.Get("pid").Val())

	data := &entity.Problem{}
	err = md.Where("pid", pid).Scan(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(data, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
