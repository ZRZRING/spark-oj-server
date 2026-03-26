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

	pid, err := md.InsertAndGetId(d)
	g.Log().Info(ctx, pid)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res.Pid = gconv.String(pid)

	return res, nil
}
