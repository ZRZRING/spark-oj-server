package contest

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type RankingReq struct {
	g.Meta    `path:"/contest/{contestId}/ranking" method:"GET" tags:"contest" summary:"获取比赛排行榜"`
	ContestId string `p:"contestId" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type RankingRes struct {
	Ranking []*RankingItem `json:"ranking" dc:"排行榜列表"`
}

type RankingItem struct {
	Rank     int                `json:"rank" dc:"排名"`
	Username string             `json:"username" dc:"用户名"`
	Score    int                `json:"score" dc:"分数"`
	Penalty  int                `json:"penalty" dc:"罚时"`
	Problems []ProblemStatsItem `json:"problems" dc:"题目状态"`
}

type ProblemStatsItem struct {
	Status      enums.RankingProblemStatus `json:"status" dc:"题目状态"`
	RejectCount int                        `json:"reject_count" dc:"错误次数"`
	FinishTime  int                        `json:"finish_time" dc:"完成时间（分钟）"`
}
