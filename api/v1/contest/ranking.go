package contest

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RankingReq struct {
	g.Meta `path:"/ranking/{cid}" method:"GET" tags:"contest" summary:"获取排行榜"`
}

type RankingRes struct{}
