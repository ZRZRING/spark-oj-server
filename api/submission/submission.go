// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package submission

import (
	"context"

	"spark-oj-server/api/submission/v1"
)

type ISubmissionV1 interface {
	GetSubmission(ctx context.Context, req *v1.GetSubmissionReq) (res *v1.GetSubmissionRes, err error)
	GetSubmissionList(ctx context.Context, req *v1.GetSubmissionListReq) (res *v1.GetSubmissionListRes, err error)
	Submit(ctx context.Context, req *v1.SubmitReq) (res *v1.SubmitRes, err error)
}
