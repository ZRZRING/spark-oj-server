package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/content"
)

func (c *ControllerContent) Create(ctx context.Context, req *content.CreateReq) (res *content.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
