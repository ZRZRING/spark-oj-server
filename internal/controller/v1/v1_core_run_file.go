package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/core"
)

func (c *ControllerCore) RunFile(ctx context.Context, req *core.RunFileReq) (res *core.RunFileRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
