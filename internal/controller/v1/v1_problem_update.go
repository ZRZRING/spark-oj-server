package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/do"

	"spark-oj/api/v1/problem"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerProblem) Update(ctx context.Context, req *problem.UpdateReq) (res *problem.UpdateRes, err error) {
	res = &problem.UpdateRes{}
	md := dao.Problem.Ctx(ctx)

	d := &do.Problem{}
	err = gconv.Struct(req, d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	problemId := gconv.Int(req.ProblemId)
	msg, err := md.Data(d).
		FieldsEx("problem_id").
		Where("problem_id", problemId).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return res, nil
}
