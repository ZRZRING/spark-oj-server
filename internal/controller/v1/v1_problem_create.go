package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/problem"
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

	problemId, err := md.InsertAndGetId(d)
	g.Log().Info(ctx, problemId)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res.ProblemId = gconv.String(problemId)

	return res, nil
}
