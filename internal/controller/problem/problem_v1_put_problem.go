package problem

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"spark-oj-server/api/problem/v1"
	"spark-oj-server/internal/consts"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"
)

func (c *ControllerV1) PutProblem(ctx context.Context, req *v1.PutProblemReq) (res *v1.PutProblemRes, err error) {
	md := dao.Problem.Ctx(ctx)
	r := g.RequestFromCtx(ctx)
	pid := gconv.String(r.Get("pid").Val())
	cnt, err := md.Where("pid", pid).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if cnt == 0 {
		err = consts.ErrProblemNotExist
		g.Log().Error(ctx, err)
	}
	data := &do.Problem{}
	err = gconv.Struct(req, &data)
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
	return
}
