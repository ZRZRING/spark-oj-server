package v1

import (
	"cmp"
	"context"
	"encoding/json"
	"slices"
	"spark-oj-server/api/v1/contest"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/pkg/consts"
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerContest) Ranking(ctx context.Context, req *contest.RankingReq) (res *contest.RankingRes, err error) {
	res = &contest.RankingRes{}
	// 获取比赛信息
	var contestInfo entity.Contest
	err = dao.Contest.Ctx(ctx).
		Where("contestId", req.ContestId).
		Scan(&contestInfo)
	if err != nil || contestInfo.ContestId == 0 {
		g.Log().Infof(ctx, "比赛不存在: %v", req.ContestId)
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "比赛不存在")
	}
	// 解析比赛题目
	var problems []int
	err = json.Unmarshal([]byte(contestInfo.Problems), &problems)
	if err != nil {
		g.Log().Infof(ctx, "解析比赛题目失败: %v", err)
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "解析比赛题目失败")
	}
	// ContestId 和 题目顺序绑定
	order := make(map[int]int)
	for i, problemId := range problems {
		order[problemId] = i
	}
	// 获取提交记录
	var total int
	submissionData := make([]entity.Submission, 0)
	err = dao.Submission.Ctx(ctx).
		Where("contestId", req.ContestId).
		OrderAsc("submissionId").
		ScanAndCount(&submissionData, &total, false)
	if err != nil {
		g.Log().Infof(ctx, "获取比赛提交列表失败: %v", err)
		return nil, gerror.Wrap(err, "获取比赛提交列表失败")
	}
	// 按 Username 计算成绩
	ranking := make(map[string]*contest.RankingItem)
	for _, sub := range submissionData {
		index := order[sub.ProblemId]
		rankingRow := ranking[sub.Username]
		if rankingRow == nil {
			rankingRow = &contest.RankingItem{
				Username: sub.Username,
				Score:    0,
				Penalty:  0,
				Problems: make([]contest.ProblemStatsItem, len(problems)),
			}
			ranking[sub.Username] = rankingRow
		}
		if rankingRow.Problems[index].Status != enums.RankingStatusAccepted {
			if sub.Result == string(enums.JudgeStatusAccepted) {
				rankingRow.Score++
				finishTime := int(sub.CreateAt.Sub(contestInfo.StartTime).Minutes())
				rankingRow.Problems[index].FinishTime = finishTime
				rankingRow.Penalty += finishTime + consts.DEFAULT_PENALTY*rankingRow.Problems[index].RejectCount
				rankingRow.Problems[index].Status = enums.RankingStatusAccepted
			} else {
				rankingRow.Problems[index].RejectCount++
				rankingRow.Problems[index].Status = enums.RankingStatusReject
			}
		}
	}
	// 丢给结果并排序
	res.Ranking = make([]*contest.RankingItem, 0, len(ranking))
	for _, it := range ranking {
		res.Ranking = append(res.Ranking, it)
	}
	slices.SortFunc(res.Ranking, func(a, b *contest.RankingItem) int {
		if n := cmp.Compare(b.Score, a.Score); n != 0 {
			return n
		}
		return cmp.Compare(a.Penalty, b.Penalty)
	})
	return res, nil
}
