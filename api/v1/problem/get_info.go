package problem

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/problem/{pid}" method:"GET" tags:"problem" summary:"获取题目信息"`
	Pid    string `p:"pid" v:"required#题目 ID 不能为空" dc:"题目 ID"`
}

type GetInfoRes struct {
	Pid         string `json:"pid" dc:""`
	Title       string `json:"title" dc:""`
	JudgeType   int    `json:"judge_type" dc:""`
	TimeLimit   int    `json:"time_limit" dc:""`
	MemoryLimit int    `json:"memory_limit" dc:""`
	Rating      int    `json:"rating" dc:""`
	CreateBy    string `json:"create_by" dc:""`
	Content     string `json:"content" dc:""`
}
