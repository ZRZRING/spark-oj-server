package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/upload"
)

func (c *ControllerUpload) TestCase(ctx context.Context, req *upload.TestCaseReq) (res *upload.TestCaseRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
