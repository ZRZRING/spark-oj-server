package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/content"
)

func (c *ControllerContent) Get(ctx context.Context, req *content.GetReq) (res *content.GetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
