package contest

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type GetProblemsReq struct {
	g.Meta    `path:"/contest/{contestId}/problems" method:"GET" tags:"contest" summary:"获取比赛题目"`
	ContestId string `p:"contest_id" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type GetProblemsRes struct {
	Total    int               `json:"total"`
	Problems []GetProblemsItem `json:"problems"`
}

type GetProblemsItem struct {
	ProblemId string          `json:"problem_id"`
	Title     string          `json:"title"`
	JudgeType enums.JudgeType `json:"judge_type"`
}
