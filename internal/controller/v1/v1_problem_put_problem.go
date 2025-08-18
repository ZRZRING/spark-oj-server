package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/problem"
)

func (c *ControllerProblem) PutProblem(ctx context.Context, req *problem.PutProblemReq) (res *problem.PutProblemRes, err error) {
	res = &problem.PutProblemRes{}
	md := dao.Problem.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	pid := gconv.String(r.Get("pid").Val())

	data := &do.Problem{}
	err = gconv.Struct(req, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	data.Pid = pid
	msg, err := md.OnConflict("pid").Save(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return res, nil
}
