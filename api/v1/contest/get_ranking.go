package contest

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type RankingReq struct {
	g.Meta `path:"/contest/{cid}/ranking/acm" method:"GET" tags:"contest" summary:"获取排行榜"`
	Cid    string `p:"cid" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type RankingRes struct {
	Item []*RankingItem `json:"item" dc:"排行榜列表"`
}

type RankingItem struct {
	Score   int                                `json:"score" dc:"分数"`
	Penalty int                                `json:"penalty" dc:"罚时"`
	Status  map[string]enums.JudgeResultStatus `json:"status" dc:"题目状态"`
}
