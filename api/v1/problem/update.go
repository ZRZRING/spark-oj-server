package problem

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta      `path:"/problem/{pid}" method:"PUT" tags:"problem" summary:"修改题目"`
	Pid         string `p:"pid" v:"required#题目 ID 不能为空" dc:"题目 ID"`
	Title       string `p:"title" v:"required#标题不能为空"`
	JudgeType   string `p:"type" v:"required#题目类型不能为空"`
	TimeLimit   int    `p:"time" v:"required#时间限制不能为空"`
	MemoryLimit int    `p:"memory" v:"required#内存限制不能为空"`
	Rating      int    `p:"rating"`
	Content     string `p:"content"`
}

type UpdateRes struct{}
