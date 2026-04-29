package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/contest"
)

func (c *ControllerContest) GetSubmissions(ctx context.Context, req *contest.GetSubmissionsReq) (res *contest.GetSubmissionsRes, err error) {
	res = &contest.GetSubmissionsRes{}
	// 判断比赛是否存在
	tot, err := dao.Contest.Ctx(ctx).Count("contest_id", req.ContestId)
	if err != nil || tot == 0 {
		g.Log().Errorf(ctx, "比赛不存在: %v", req.ContestId)
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "比赛不存在")
	}
	// 获取比赛的所有提交
	submissionData := make([]*entity.Submission, 0)
	err = dao.Submission.Ctx(ctx).
		Where("contest_id", req.ContestId).
		Page(req.Page, req.Size).
		OrderDesc("submission_id").
		ScanAndCount(&submissionData, &res.Total, false)
	if err != nil {
		g.Log().Errorf(ctx, "获取比赛提交列表失败: %v", err)
		return nil, gerror.Wrap(err, "获取比赛提交列表失败")
	}
	// 绑定回结果
	for _, sub := range submissionData {
		res.Submissions = append(res.Submissions, &contest.GetSubmissionsItem{
			SubmissionId: gconv.String(sub.SubmissionId),
			ProblemId:    gconv.String(sub.ProblemId),
			ContestId:    gconv.String(sub.ContestId),
			Username:     sub.Username,
			Result:       sub.Result,
			TimeCost:     gconv.String(sub.TimeCost),
			MemoryCost:   gconv.String(sub.MemoryCost),
			Language:     sub.Language,
			CreateTime:   sub.CreateAt.Layout("2006-01-02 15:04:05"),
		})
	}
	return res, nil
}
