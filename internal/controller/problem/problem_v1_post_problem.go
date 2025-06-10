package problem

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"spoj/api/problem/v1"
	"spoj/internal/consts"
	"spoj/internal/dao"
	"spoj/internal/model/do"
)

func (c *ControllerV1) PostProblem(ctx context.Context, req *v1.PostProblemReq) (res *v1.PostProblemRes, err error) {
	g.Log().Info(ctx, req)

	// 校验参数
	err = g.Validator().Data(req).Run(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 检查题目是否存在
	md := dao.Problem.Ctx(ctx)
	cnt, err := md.Where("pid", req.Pid).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if cnt == 0 {
		err = consts.ErrProblemNotExist
		g.Log().Error(ctx, err)
	}

	// 往数据库添加题目
	msg, err := md.Insert(do.Problem{
		Pid:         req.Pid,
		Title:       req.Title,
		JudgeType:   req.JudgeType,
		Time:        req.Time,
		Memory:      req.Memory,
		ContentType: req.ContentType,
		Content:     req.Content,
		CreateBy:    req.CreateBy,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return
}
