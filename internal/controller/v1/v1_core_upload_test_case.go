package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/core"
)

func (c *ControllerCore) UploadTestCase(ctx context.Context, req *core.UploadTestCaseReq) (res *core.UploadTestCaseRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
