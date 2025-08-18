package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/submission"
)

func (c *ControllerSubmission) GetSubmissionList(ctx context.Context, req *submission.GetSubmissionListReq) (res *submission.GetSubmissionListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
