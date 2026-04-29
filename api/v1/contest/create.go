package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta      `path:"/contest" method:"POST" tags:"contest" summary:"添加比赛"`
	Title       string     `json:"title" v:"required#比赛名称不能为空"`
	Password    string     `json:"password"`
	Problems    []int      `json:"problems"`
	Description string     `json:"description"`
	Practice    bool       `json:"practice"`
	StartTime   gtime.Time `json:"start_time" v:"required-if:practice,false#开始时间不能为空"`
	EndTime     gtime.Time `json:"end_time" v:"required-if:practice,false#结束时间不能为空"`
	CreateBy    string     `json:"create_by" v:"required#创建者不能为空"`
	LockTime    gtime.Time `json:"lock_time" v:"after:start_time|before:end_time#封榜时间不合法"`
}

type CreateRes struct {
}
