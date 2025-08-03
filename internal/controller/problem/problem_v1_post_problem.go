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
	res = &v1.PostProblemRes{}
	md := dao.Problem.Ctx(ctx)

	data := &do.Problem{}
	err = gconv.Struct(req, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	msg, err := md.Insert(data)
	g.Log().Info(ctx, msg)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
