package file

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SubmitReq struct {
	g.Meta   `path:"/submission" method:"POST" tag:"submission" summary:"提交题目"`
	Pid      string `json:"pid" v:"required#题目编号不能为空"`
	Username string `json:"username" v:"required#提交用户不能为空"`
	Cid      string `json:"cid"`
	Language string `json:"language" default:"cpp"`
	Code     string `json:"code"`
}

type SubmitRes struct {
	Sid    int    `json:"sid"`
	Result string `json:"result"`
}
