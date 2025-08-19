package contest

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/contest/{cid}" method:"GET" tags:"contest" summary:"获取比赛信息"`
}

type GetRes struct {
	Title       string `p:"title" dc:""`
	Problems    any    `p:"problems" dc:""`
	Description string `p:"description" dc:""`
	StartTime   int64  `p:"start_time" dc:""`
	EndTime     int64  `p:"end_time" dc:""`
	LockTime    int64  `p:"lock_time" dc:""`
	CreateBy    string `p:"create_by" dc:""`
}
