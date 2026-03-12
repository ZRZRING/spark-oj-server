package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/submission"
)

func (c *ControllerSubmission) GetPage(ctx context.Context, req *submission.GetPageReq) (res *submission.GetPageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
