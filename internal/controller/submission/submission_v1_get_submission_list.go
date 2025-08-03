package submission

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/submission/v1"
)

func (c *ControllerV1) GetSubmissionList(ctx context.Context, req *v1.GetSubmissionListReq) (res *v1.GetSubmissionListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
