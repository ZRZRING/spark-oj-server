package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProblemReq struct {
	g.Meta `path:"/problem/{pid}" method:"GET" tags:"problem" summary:"获取题目信息"`
}

type GetProblemRes struct {
	Title     string `json:"title" dc:""`
	JudgeType int    `json:"judge_type" dc:""`
	Time      int    `json:"time" dc:""`
	Memory    int    `json:"memory" dc:""`
	Content   string `json:"content" dc:""`
	CreateBy  string `json:"create_by" dc:""`
}
