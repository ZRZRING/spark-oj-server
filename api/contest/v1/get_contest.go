package v1

import "github.com/gogf/gf/v2/frame/g"

type GetContestReq struct {
	g.Meta `path:"/contest/{cid}" method:"GET" tags:"contest" summary:"获取比赛信息"`
}

type GetContestRes struct {
	Title       string `json:"title" dc:""`
	Problems    any    `json:"problems" dc:""`
	Description string `json:"description" dc:""`
	StartTime   int64  `json:"start_time" dc:""`
	EndTime     int64  `json:"end_time" dc:""`
	LockTime    int64  `json:"lock_time" dc:""`
	CreateBy    string `json:"create_by" dc:""`
}
