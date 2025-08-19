// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"spark-oj-server/api/v1/admin"
	"spark-oj-server/api/v1/contest"
	"spark-oj-server/api/v1/core"
	"spark-oj-server/api/v1/problem"
	"spark-oj-server/api/v1/submission"
	"spark-oj-server/api/v1/upload"
	"spark-oj-server/api/v1/user"
)

type IV1Admin interface {
	Protected(ctx context.Context, req *admin.ProtectedReq) (res *admin.ProtectedRes, err error)
}

type IV1Contest interface {
	Create(ctx context.Context, req *contest.CreateReq) (res *contest.CreateRes, err error)
	Get(ctx context.Context, req *contest.GetReq) (res *contest.GetRes, err error)
	GetList(ctx context.Context, req *contest.GetListReq) (res *contest.GetListRes, err error)
	Update(ctx context.Context, req *contest.UpdateReq) (res *contest.UpdateRes, err error)
}

type IV1Core interface {
	Judge(ctx context.Context, req *core.JudgeReq) (res *core.JudgeRes, err error)
	Run(ctx context.Context, req *core.RunReq) (res *core.RunRes, err error)
}

type IV1Problem interface {
	Create(ctx context.Context, req *problem.CreateReq) (res *problem.CreateRes, err error)
	Get(ctx context.Context, req *problem.GetReq) (res *problem.GetRes, err error)
	GetList(ctx context.Context, req *problem.GetListReq) (res *problem.GetListRes, err error)
	Update(ctx context.Context, req *problem.UpdateReq) (res *problem.UpdateRes, err error)
}

type IV1Submission interface {
	Get(ctx context.Context, req *submission.GetReq) (res *submission.GetRes, err error)
	GetList(ctx context.Context, req *submission.GetListReq) (res *submission.GetListRes, err error)
}

type IV1Upload interface {
	TestCase(ctx context.Context, req *upload.TestCaseReq) (res *upload.TestCaseRes, err error)
}

type IV1User interface {
	Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error)
	Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error)
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
	Training(ctx context.Context, req *user.TrainingReq) (res *user.TrainingRes, err error)
}
