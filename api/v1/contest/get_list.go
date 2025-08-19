package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetListReq struct {
	g.Meta `path:"/contests" method:"GET" summary:"获取比赛列表" tags:"contest"`
	Page   int `p:"page" v:"required#页码不能为空" dc:""`
	Size   int `p:"size" v:"required#每页数量不能为空" dc:""`
}

type GetListRes struct {
	Total    int `json:"total"`
	Contests []*struct {
		Cid       string     `json:"cid"`
		Title     string     `json:"title"`
		Practice  bool       `json:"practice" dc:"true 表示改比赛不限制参与时间，为训练题单"`
		StartTime gtime.Time `json:"start_time"`
		EndTime   gtime.Time `json:"end_time"`
	} `json:"contests" dc:"比赛列表"`
}
