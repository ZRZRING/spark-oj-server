package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetInfoReq struct {
	g.Meta `path:"/contest/{cid}" method:"GET" tags:"contest" summary:"获取比赛信息"`
	Cid    string `p:"cid" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type GetInfoRes struct {
	Cid         string     `json:"cid" dc:""`
	Title       string     `json:"title" dc:""`
	Description string     `json:"description" dc:""`
	StartTime   gtime.Time `json:"start_time" dc:""`
	EndTime     gtime.Time `json:"end_time" dc:""`
	LockTime    gtime.Time `json:"lock_time" dc:""`
	CreateBy    string     `json:"create_by" dc:""`
}
