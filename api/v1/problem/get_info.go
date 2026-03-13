package problem

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/problem/{pid}" method:"GET" tags:"problem" summary:"获取题目信息"`
	Pid    string `p:"pid" v:"required#题目 ID 不能为空" dc:"题目 ID"`
}

type GetInfoRes struct {
	Pid         string `json:"pid"`
	Title       string `json:"title"`
	JudgeType   int    `json:"judge_type"`
	TimeLimit   int    `json:"time_limit"`
	MemoryLimit int    `json:"memory_limit"`
	Rating      int    `json:"rating"`
	CreateBy    string `json:"create_by"`
	Content     string `json:"content"`
}
