package contest

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetPageReq struct {
	g.Meta `path:"/contests" method:"GET" summary:"获取比赛列表" tags:"contest"`
	Page   int `p:"page" v:"required#页码不能为空"`
	Size   int `p:"size" v:"required#每页数量不能为空"`
}

type GetPageRes struct {
	Total    int `json:"total"`
	Contests []*struct {
		ContestId  string                  `json:"contest_id"`
		Title      string                  `json:"title"`
		Practice   bool                    `json:"practice" dc:"true 表示比赛不限制参与时间，为训练题单"`
		StartTime  gtime.Time              `json:"start_time"`
		EndTime    gtime.Time              `json:"end_time"`
		CreateBy   string                  `json:"create_by"`
		Visibility enums.ContestVisibility `json:"visibility"`
	} `json:"contests" dc:"比赛列表"`
}
