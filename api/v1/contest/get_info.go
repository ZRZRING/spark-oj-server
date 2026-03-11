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
	Cid         string     `p:"cid" dc:""`
	Title       string     `p:"title" dc:""`
	Description string     `p:"description" dc:""`
	StartTime   gtime.Time `p:"start_time" dc:""`
	EndTime     gtime.Time `p:"end_time" dc:""`
	LockTime    gtime.Time `p:"lock_time" dc:""`
	CreateBy    string     `p:"create_by" dc:""`
}
