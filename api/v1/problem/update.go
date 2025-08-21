package problem

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta      `path:"/problem/{pid}" method:"PUT" tags:"problem" mime:"application/json" summary:"修改题目"`
	Title       string `p:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `p:"type" v:"required#题目类型不能为空" dc:""`
	TimeLimit   int    `p:"time" v:"required#时间限制不能为空" dc:""`
	MemoryLimit int    `p:"memory" v:"required#内存限制不能为空" dc:""`
}

type UpdateRes struct{}
