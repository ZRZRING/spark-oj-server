package problem

import "github.com/gogf/gf/v2/frame/g"

type PutProblemReq struct {
	g.Meta      `path:"/problem/{pid}" method:"put" tags:"problem" summary:"修改题目"`
	Title       string `json:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `json:"type" v:"required#题目类型不能为空" dc:""`
	TimeLimit   int    `json:"time" v:"required#时间限制不能为空" dc:""`
	MemoryLimit int    `json:"memory" v:"required#内存限制不能为空" dc:""`
}

type PutProblemRes struct{}
