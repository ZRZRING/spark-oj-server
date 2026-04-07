// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"spark-oj/api/v1/admin"
	"spark-oj/api/v1/contest"
	"spark-oj/api/v1/core"
	"spark-oj/api/v1/problem"
	"spark-oj/api/v1/submission"
	"spark-oj/api/v1/user"
)

type IV1Admin interface {
	Protected(ctx context.Context, req *admin.ProtectedReq) (res *admin.ProtectedRes, err error)
}

type IV1Contest interface {
	Create(ctx context.Context, req *contest.CreateReq) (res *contest.CreateRes, err error)
	GetInfo(ctx context.Context, req *contest.GetInfoReq) (res *contest.GetInfoRes, err error)
	GetPage(ctx context.Context, req *contest.GetPageReq) (res *contest.GetPageRes, err error)
	GetProblemInfo(ctx context.Context, req *contest.GetProblemInfoReq) (res *contest.GetProblemInfoRes, err error)
	GetProblems(ctx context.Context, req *contest.GetProblemsReq) (res *contest.GetProblemsRes, err error)
	Ranking(ctx context.Context, req *contest.RankingReq) (res *contest.RankingRes, err error)
	GetSubmissions(ctx context.Context, req *contest.GetSubmissionsReq) (res *contest.GetSubmissionsRes, err error)
	Update(ctx context.Context, req *contest.UpdateReq) (res *contest.UpdateRes, err error)
}

type IV1Core interface {
	Judge(ctx context.Context, req *core.JudgeReq) (res *core.JudgeRes, err error)
	Run(ctx context.Context, req *core.RunReq) (res *core.RunRes, err error)
	UploadTestCase(ctx context.Context, req *core.UploadTestCaseReq) (res *core.UploadTestCaseRes, err error)
}

type IV1Problem interface {
	Create(ctx context.Context, req *problem.CreateReq) (res *problem.CreateRes, err error)
	GetInfo(ctx context.Context, req *problem.GetInfoReq) (res *problem.GetInfoRes, err error)
	GetPage(ctx context.Context, req *problem.GetPageReq) (res *problem.GetPageRes, err error)
	Update(ctx context.Context, req *problem.UpdateReq) (res *problem.UpdateRes, err error)
}

type IV1Submission interface {
	GetInfo(ctx context.Context, req *submission.GetInfoReq) (res *submission.GetInfoRes, err error)
	GetPage(ctx context.Context, req *submission.GetPageReq) (res *submission.GetPageRes, err error)
}

type IV1User interface {
	GetPage(ctx context.Context, req *user.GetPageReq) (res *user.GetPageRes, err error)
	Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error)
	Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error)
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
	Training(ctx context.Context, req *user.TrainingReq) (res *user.TrainingRes, err error)
}
