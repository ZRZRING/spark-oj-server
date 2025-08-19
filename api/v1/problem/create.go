package problem

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta      `path:"/problem" method:"POST" tags:"problem" summary:"添加题目"`
	Pid         string `p:"pid" v:"required#题目ID不能为空" dc:""`
	Title       string `p:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `p:"judge_type" v:"required#判题类型不能为空" dc:""`
	TimeLimit   int    `p:"time_limit" v:"required#时间限制不能为空" dc:""`
	MemoryLimit int    `p:"memory_limit" v:"required#内存限制不能为空" dc:""`
	CreateBy    string `p:"create_by" v:"required#创建者不能为空" dc:""`
}

type CreateRes struct{}
