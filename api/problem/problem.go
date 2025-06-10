// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package problem

import (
	"context"

	"spoj/api/problem/v1"
)

type IProblemV1 interface {
	GetProblem(ctx context.Context, req *v1.GetProblemReq) (res *v1.GetProblemRes, err error)
	GetProblemList(ctx context.Context, req *v1.GetProblemListReq) (res *v1.GetProblemListRes, err error)
	PostProblem(ctx context.Context, req *v1.PostProblemReq) (res *v1.PostProblemRes, err error)
	PutProblem(ctx context.Context, req *v1.PutProblemReq) (res *v1.PutProblemRes, err error)
}
