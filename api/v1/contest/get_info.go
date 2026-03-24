package contest

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta `path:"/contest/{cid}" method:"GET" tags:"contest" summary:"获取比赛信息"`
	Cid    string `p:"cid" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type GetInfoRes struct {
	Cid         string `json:"cid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	LockTime    string `json:"lock_time"`
	CreateBy    string `json:"create_by"`
}
