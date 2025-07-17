// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package contest

import (
	"context"

	"spark-oj-server/api/contest/v1"
)

type IContestV1 interface {
	GetContest(ctx context.Context, req *v1.GetContestReq) (res *v1.GetContestRes, err error)
	GetContestList(ctx context.Context, req *v1.GetContestListReq) (res *v1.GetContestListRes, err error)
	PostContest(ctx context.Context, req *v1.PostContestReq) (res *v1.PostContestRes, err error)
	PutContest(ctx context.Context, req *v1.PutContestReq) (res *v1.PutContestRes, err error)
}
