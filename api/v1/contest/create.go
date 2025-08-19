package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta       `path:"/contest" method:"POST" tags:"contest" summary:"添加比赛"`
	Cid          string     `p:"cid" v:"required#比赛ID不能为空" dc:""`
	Title        string     `p:"title" v:"required#比赛名称不能为空" dc:""`
	Password     string     `p:"password" dc:""`
	Problems     g.Map      `p:"problems" dc:""`
	Description  string     `p:"description" dc:""`
	TimeRequired bool       `p:"time_required" dc:""`
	StartTime    gtime.Time `p:"start_time" v:"required-if:timeRequired#开始时间不能为空" dc:""`
	EndTime      gtime.Time `p:"end_time" v:"required-if:timeRequired#结束时间不能为空" dc:""`
	CreateBy     string     `p:"create_by" v:"required#创建者不能为空" dc:""`
	LockTime     gtime.Time `p:"lock_time" v:"between:startTime,endTime#封榜时间不合法" dc:""`
}

type CreateRes struct {
}
