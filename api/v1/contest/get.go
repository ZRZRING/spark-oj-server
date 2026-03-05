package contest

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/contest/{cid}" method:"GET" tags:"contest" summary:"获取比赛信息"`
	Cid    string `p:"cid" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
}

type GetRes struct {
	Cid         string `p:"cid" dc:""`
	Title       string `p:"title" dc:""`
	Problems    any    `p:"problems" dc:""`
	Description string `p:"description" dc:""`
	StartTime   int64  `p:"start_time" dc:""`
	EndTime     int64  `p:"end_time" dc:""`
	LockTime    int64  `p:"lock_time" dc:""`
	CreateBy    string `p:"create_by" dc:""`
}
