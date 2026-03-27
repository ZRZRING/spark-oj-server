package v1

import (
	"context"
	"encoding/json"
	"spark-oj-server/api/v1/contest"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/pkg/consts"
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerContest) Ranking(ctx context.Context, req *contest.RankingReq) (res *contest.RankingRes, err error) {
	res = &contest.RankingRes{}
	// 获取比赛题目
	var data string
	err = dao.Contest.Ctx(ctx).
		Where("cid", req.Cid).
		Fields("problems").
		Scan(&data)
	if err != nil || data == "" {
		g.Log().Infof(ctx, "比赛不存在或题目不存在: %v", req.Cid)
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "比赛不存在")
	}
	// 解析比赛题目
	var problems []string
	err = json.Unmarshal([]byte(data), &problems)
	if err != nil {
		g.Log().Infof(ctx, "解析比赛信息失败: %v", err)
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "解析比赛信息失败")
	}
	// Cid 和 题目顺序绑定
	order := make(map[string]int)
	for i, pid := range problems {
		order[pid] = i
	}
	// 获取提交记录
	var total int
	submissionData := make([]entity.Submission, 0)
	err = dao.Submission.Ctx(ctx).
		Where("cid", req.Cid).
		OrderAsc("sid").
		ScanAndCount(&submissionData, &total, false)
	if err != nil {
		g.Log().Infof(ctx, "获取比赛提交列表失败: %v", err)
		return nil, gerror.Wrap(err, "获取比赛提交列表失败")
	}
	// 按 Username 计算成绩
	Ranking := make(map[string]*contest.RankingItem)
	for _, submission := range submissionData {
		pid := gconv.String(submission.Pid)
		it := Ranking[submission.Username]
		if it == nil {
			it = &contest.RankingItem{Username: submission.Username}
			Ranking[submission.Username] = it
		}
		if it.Status[pid] != enums.RankingStatusAccepted {
			if submission.Result == string(enums.JudgeStatusAccepted) {
				it.Score++
				it.Status[pid] = enums.RankingStatusAccepted
			} else {
				it.Penalty += consts.DEFAULT_PENALTY
				it.Status[pid] = enums.RankingStatusReject
			}
		}
	}
	return res, nil
}
