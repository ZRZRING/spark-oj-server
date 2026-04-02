package problem

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta      `path:"/problem" method:"POST" tags:"problem" summary:"添加题目"`
	Title       string          `p:"title" v:"required#标题不能为空"`
	JudgeType   enums.JudgeType `p:"judge_type" v:"required#判题类型不能为空"`
	TimeLimit   int             `p:"time_limit" v:"required#时间限制不能为空"`
	MemoryLimit int             `p:"memory_limit" v:"required#内存限制不能为空"`
	Rating      int             `p:"rating"`
	CreateBy    string          `p:"create_by" v:"required#创建者不能为空"`
	Content     string          `p:"content"`
}

type CreateRes struct {
	ProblemId string `json:"problemId"`
}
