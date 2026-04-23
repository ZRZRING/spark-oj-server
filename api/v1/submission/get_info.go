package submission

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta       `path:"/submission/{submissionId}" method:"GET" tags:"submission" summary:"获取提交信息"`
	SubmissionId string `p:"submissionId" in:"path" v:"required#提交 ID 不能为空"`
}

type GetInfoRes struct {
	SubmissionId string `json:"submission_id"`
	ProblemId    string `json:"problem_id"`
	ContestId    string `json:"contest_id"`
	Username     string `json:"username"`
	Result       string `json:"result"`
	Language     string `json:"language"`
	MemoryCost   string `json:"memory_cost"`
	TimeCost     string `json:"time_cost"`
	Code         string `json:"code"`
	CreateTime   string `json:"create_time"`
}
