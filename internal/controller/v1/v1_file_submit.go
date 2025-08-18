package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/file"
)

func (c *ControllerFile) Submit(ctx context.Context, req *file.SubmitReq) (res *file.SubmitRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
