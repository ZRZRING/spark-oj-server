package core

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type JudgeReq struct {
	g.Meta    `path:"/judge" method:"POST" tags:"core" summary:"提交题目"`
	ProblemId string         `json:"problemId" v:"required"`
	ContestId string         `json:"contestId"`
	Username  string         `json:"username" v:"required"`
	Code      string         `json:"code" v:"required"`
	Language  enums.Language `json:"language" v:"required" default:"cpp"`
}

type JudgeRes struct {
	SubmissionId string            `json:"submissionId"`
	Result       enums.JudgeStatus `json:"result"`
}
