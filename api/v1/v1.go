// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"spark-oj-server/api/v1/admin"
	"spark-oj-server/api/v1/contest"
	"spark-oj-server/api/v1/file"
	"spark-oj-server/api/v1/problem"
	"spark-oj-server/api/v1/submission"
	"spark-oj-server/api/v1/user"
)

type IV1Admin interface {
	Protected(ctx context.Context, req *admin.ProtectedReq) (res *admin.ProtectedRes, err error)
}

type IV1Contest interface {
	GetContest(ctx context.Context, req *contest.GetContestReq) (res *contest.GetContestRes, err error)
	GetContestList(ctx context.Context, req *contest.GetContestListReq) (res *contest.GetContestListRes, err error)
	PostContest(ctx context.Context, req *contest.PostContestReq) (res *contest.PostContestRes, err error)
	PutContest(ctx context.Context, req *contest.PutContestReq) (res *contest.PutContestRes, err error)
}

type IV1File interface {
	Submit(ctx context.Context, req *file.SubmitReq) (res *file.SubmitRes, err error)
}

type IV1Problem interface {
	GetProblem(ctx context.Context, req *problem.GetProblemReq) (res *problem.GetProblemRes, err error)
	GetProblemList(ctx context.Context, req *problem.GetProblemListReq) (res *problem.GetProblemListRes, err error)
	PostProblem(ctx context.Context, req *problem.PostProblemReq) (res *problem.PostProblemRes, err error)
	PutProblem(ctx context.Context, req *problem.PutProblemReq) (res *problem.PutProblemRes, err error)
}

type IV1Submission interface {
	GetSubmission(ctx context.Context, req *submission.GetSubmissionReq) (res *submission.GetSubmissionRes, err error)
	GetSubmissionList(ctx context.Context, req *submission.GetSubmissionListReq) (res *submission.GetSubmissionListRes, err error)
}

type IV1User interface {
	Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error)
	Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error)
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
	Training(ctx context.Context, req *user.TrainingReq) (res *user.TrainingRes, err error)
}
