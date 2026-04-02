package submission

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta       `path:"/submission/{submissionId}" method:"GET" tags:"submission" summary:"获取提交信息"`
	SubmissionId string `p:"submissionId" v:"required#题目 ID 不能为空"`
}

type GetInfoRes struct {
	SubmissionId string `json:"submissionId"`
	ProblemId    string `json:"problemId"`
	ContestId    string `json:"contestId"`
	Username     string `json:"username"`
	Result       string `json:"result"`
	Language     string `json:"language"`
	MemoryCost   string `json:"memory_cost"`
	TimeCost     string `json:"time_cost"`
	Code         string `json:"code"`
	CreateAt     string `json:"create_at"`
}
