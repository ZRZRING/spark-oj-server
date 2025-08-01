package problem

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/problem/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
)

func (c *ControllerV1) GetProblem(ctx context.Context, req *v1.GetProblemReq) (res *v1.GetProblemRes, err error) {
	res = &v1.GetProblemRes{}
	md := dao.Problem.Ctx(ctx)

	// 绑定 URL 信息
	r := g.RequestFromCtx(ctx)
	pid := gconv.String(r.Get("pid").Val())

	// 从数据库获取数据
	problem := &entity.Problem{}
	err = md.Where("pid", pid).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 处理返回信息
	err = gconv.Scan(problem, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
