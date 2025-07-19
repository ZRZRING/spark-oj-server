package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProblemReq struct {
	g.Meta `path:"/problem/{pid}" method:"GET" tags:"problem" summary:"获取题目信息"`
}

type GetProblemRes struct {
	Title       string `json:"title" dc:""`
	JudgeType   int    `json:"judge_type" dc:""`
	TimeLimit   int    `json:"time_limit" dc:""`
	MemoryLimit int    `json:"memory_limit" dc:""`
	CreateBy    string `json:"create_by" dc:""`
}
