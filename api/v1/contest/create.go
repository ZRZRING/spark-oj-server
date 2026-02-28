package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta      `path:"/contest" method:"POST" tags:"contest" mime:"application/json" summary:"添加比赛"`
	Title       string     `json:"title" v:"required#比赛名称不能为空" dc:""`
	Password    string     `json:"password" dc:""`
	Problems    []int      `json:"problems" dc:""`
	Description string     `json:"description" dc:""`
	Practice    bool       `json:"practice" dc:""`
	StartTime   gtime.Time `json:"start_time" v:"required-if:practice,false#开始时间不能为空" dc:""`
	EndTime     gtime.Time `json:"end_time" v:"required-if:practice,false#结束时间不能为空" dc:""`
	CreateBy    string     `json:"create_by" v:"required#创建者不能为空" dc:""`
	LockTime    gtime.Time `json:"lock_time" v:"between:startTime,endTime#封榜时间不合法" dc:""`
}

type CreateRes struct {
}
