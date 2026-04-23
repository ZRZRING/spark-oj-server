package problem

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type UpdateReq struct {
	g.Meta      `path:"/problem/{problemId}" method:"PUT" tags:"problem" summary:"修改题目"`
	ProblemId   string          `p:"problemId" in:"path" v:"required#题目 ID 不能为空"`
	Title       string          `json:"title" v:"required#标题不能为空"`
	JudgeType   enums.JudgeType `json:"judge_type" v:"required#题目类型不能为空"`
	TimeLimit   int             `json:"time_limit" v:"required#时间限制不能为空"`
	MemoryLimit int             `json:"memory_limit" v:"required#内存限制不能为空"`
	Rating      int             `json:"rating"`
	Content     string          `json:"content"`
}

type UpdateRes struct{}
