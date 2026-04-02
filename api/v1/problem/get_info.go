package problem

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta    `path:"/problem/{problemId}" method:"GET" tags:"problem" summary:"获取题目信息"`
	ProblemId string `p:"problemId" v:"required#题目 ID 不能为空" dc:"题目 ID"`
}

type GetInfoRes struct {
	ProblemId   string          `json:"problemId"`
	Title       string          `json:"title"`
	JudgeType   enums.JudgeType `json:"judge_type"`
	TimeLimit   int             `json:"time_limit"`
	MemoryLimit int             `json:"memory_limit"`
	Rating      int             `json:"rating"`
	CreateBy    string          `json:"create_by"`
	Content     string          `json:"content"`
}
