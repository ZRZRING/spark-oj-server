package submission

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/submission/v1"
)

func (c *ControllerV1) Submit(ctx context.Context, req *v1.SubmitReq) (res *v1.SubmitRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
