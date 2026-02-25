package v1

import (
	"context"
	"spark-oj-server/api/v1/core"
	"spark-oj-server/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerCore) Run(ctx context.Context, req *core.RunReq) (res *core.RunRes, err error) {
	res = &core.RunRes{}

	exeReq := &service.ExecuteCodeRequest{}
	err = gconv.Struct(req, exeReq)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	exeRes, err := service.ExecuteCode(ctx, exeReq)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Struct(exeRes, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 纳秒转化成毫秒
	res.Time /= 1000 * 1000

	return res, nil
}
