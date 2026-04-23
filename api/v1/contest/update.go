package contest

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UpdateReq struct {
	g.Meta      `path:"/contest/{contestId}" method:"PUT" tags:"contest" summary:"更新比赛"`
	ContestId   string     `p:"contestId" in:"path" v:"required#比赛 ID 不能为空"`
	Title       string     `json:"title" v:"required#比赛名称不能为空"`
	Password    string     `json:"password"`
	Problems    []int      `json:"problems" dc:"比赛题目，支持数组或对象格式"`
	Description string     `json:"description"`
	Practice    bool       `json:"practice"`
	StartTime   gtime.Time `json:"start_time" v:"required-if:practice,false#开始时间不能为空"`
	EndTime     gtime.Time `json:"end_time" v:"required-if:practice,false#结束时间不能为空"`
	CreateBy    string     `json:"create_by" v:"required#创建者不能为空"`
	LockTime    gtime.Time `json:"lock_time" v:"between:startTime,endTime#封榜时间不合法"`
}

type UpdateRes struct {
}
