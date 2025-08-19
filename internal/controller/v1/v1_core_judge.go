package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/core"
)

func (c *ControllerCore) Judge(ctx context.Context, req *core.JudgeReq) (res *core.JudgeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
