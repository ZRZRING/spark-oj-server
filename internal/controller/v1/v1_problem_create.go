package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) Create(ctx context.Context, req *problem.CreateReq) (res *problem.CreateRes, err error) {
	res = &problem.CreateRes{}
	md := dao.Problem.Ctx(ctx)

	d := &do.Problem{}
	err = gconv.Struct(req, d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	msg, err := md.Insert(d)
	g.Log().Info(ctx, msg)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
