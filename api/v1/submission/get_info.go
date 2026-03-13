package submission

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta `path:"/submission/{sid}" method:"GET" tags:"submission" summary:"获取提交信息"`
}

type GetInfoRes struct {
	ProblemId  string `json:"problem_id"`
	Username   string `json:"username"`
	ContestId  string `json:"contest_id"`
	Result     string `json:"result"`
	Language   string `json:"language"`
	MemoryCost int    `json:"memory_cost"`
	TimeCost   int    `json:"time_cost"`
	Code       string `json:"code"`
	CreateAt   string `json:"create_at"`
}
