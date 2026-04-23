package problem

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta      `path:"/problem" method:"POST" tags:"problem" summary:"添加题目"`
	Title       string          `json:"title" v:"required#标题不能为空"`
	JudgeType   enums.JudgeType `json:"judge_type" v:"required#判题类型不能为空"`
	TimeLimit   int             `json:"time_limit" v:"required#时间限制不能为空"`
	MemoryLimit int             `json:"memory_limit" v:"required#内存限制不能为空"`
	Rating      int             `json:"rating"`
	Content     string          `json:"content"`
	CreateBy    string          `json:"create_by" v:"required#创建者不能为空"`
}

type CreateRes struct {
	ProblemId string `json:"problem_id"`
}
