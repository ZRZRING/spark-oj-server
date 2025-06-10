package v1

import "github.com/gogf/gf/v2/frame/g"

type PostProblemReq struct {
	g.Meta      `path:"/problem" method:"post" tags:"problem" summary:"添加题目"`
	Pid         string `json:"pid" v:"required#题目ID不能为空" dc:""`
	Title       string `json:"title" v:"required#标题不能为空" dc:""`
	JudgeType   string `json:"type" v:"required#判题类型不能为空" dc:""`
	Time        int    `json:"time" v:"required#时间限制不能为空" dc:""`
	Memory      int    `json:"memory" v:"required#内存限制不能为空" dc:""`
	ContentType string `json:"content_type" v:"required#内容类型不能为空" dc:""`
	Content     string `json:"content" dc:""`
	CreateBy    string `json:"create_by" v:"required#创建者不能为空" dc:""`
}

type PostProblemRes struct{}
