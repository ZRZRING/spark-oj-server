package submission

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSubmissionReq struct {
	g.Meta `path:"/submission/{sid}" method:"GET" tags:"contest" summary:"获取提交信息"`
}

type GetSubmissionRes struct {
	ProblemId  string `json:"problem_id" dc:""`
	Username   string `json:"username" dc:""`
	ContestId  string `json:"contest_id" dc:""`
	Result     string `json:"result" dc:""`
	Language   string `json:"language" dc:""`
	MemoryCost int    `json:"memory_cost" dc:""`
	TimeCost   int    `json:"time_cost" dc:""`
	Code       string `json:"code" dc:""`
	CreateAt   string `json:"create_at" dc:""`
}
