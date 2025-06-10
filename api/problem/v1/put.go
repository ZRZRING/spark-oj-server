package v1

import "github.com/gogf/gf/v2/frame/g"

type PutProblemReq struct {
	g.Meta      `path:"/problem/{pid}" method:"put" tags:"problem" summary:"修改题目"`
	Pid         string `json:"pid" v:"required#题目ID不能为空" dc:""`
	Title       string `json:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `json:"type" v:"required#题目类型不能为空" dc:""`
	Time        int    `json:"time" v:"required#时间限制不能为空" dc:""`
	Memory      int    `json:"memory" v:"required#内存限制不能为空" dc:""`
	ContentType string `json:"content_type" v:"required#内容类型不能为空" dc:""`
	Content     string `json:"content" dc:""`
}

type PutProblemRes struct{}
