package problem

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/problem/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PostProblem(ctx context.Context, req *v1.PostProblemReq) (res *v1.PostProblemRes, err error) {
	g.Log().Info(ctx, req)
	md := dao.Problem.Ctx(ctx)
	data := &do.Problem{}
	err = gconv.Struct(req, &data)
	msg, err := md.Insert(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)
	return
}
