package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetContestListReq struct {
	g.Meta `path:"/contests" method:"get" summary:"获取比赛列表" tags:"contest"`
	Page   int `json:"page" v:"required#页码不能为空" dc:""`
	Size   int `json:"size" v:"required#每页数量不能为空" dc:""`
}

type GetContestListRes struct {
	Contests []*Contest `json:"contests" dc:"比赛列表"`
}

type Contest struct {
	Cid       string     `json:"cid"`
	Title     string     `json:"title"`
	Practice  bool       `json:"practice" dc:"true 表示改比赛不限制参与时间，为训练题单"`
	StartTime gtime.Time `json:"start_time"`
	EndTime   gtime.Time `json:"end_time"`
}
