package problem

import "github.com/gogf/gf/v2/frame/g"

type PostProblemReq struct {
	g.Meta      `path:"/problem" method:"post" tags:"problem" summary:"添加题目"`
	Pid         string `json:"pid" v:"required#题目ID不能为空" dc:""`
	Title       string `json:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `json:"judge_type" v:"required#判题类型不能为空" dc:""`
	TimeLimit   int    `json:"time_limit" v:"required#时间限制不能为空" dc:""`
	MemoryLimit int    `json:"memory_limit" v:"required#内存限制不能为空" dc:""`
	CreateBy    string `json:"create_by" v:"required#创建者不能为空" dc:""`
}

type PostProblemRes struct{}
