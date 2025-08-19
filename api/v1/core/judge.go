package core

import (
	"github.com/gogf/gf/v2/frame/g"
)

type JudgeReq struct {
	g.Meta   `path:"/submission" method:"POST" tag:"submission" summary:"提交题目"`
	Pid      string `p:"pid" v:"required"`
	Code     string `p:"code" v:"required"`
	Username string `p:"username" v:"required"`
	Language string `p:"language" default:"cpp"`
	Cid      string `p:"cid"`
}

type JudgeRes struct {
	Sid    int    `json:"sid"`
	Result string `json:"result"`
}
