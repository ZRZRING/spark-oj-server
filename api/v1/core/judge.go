package core

import (
	"github.com/gogf/gf/v2/frame/g"
)

type JudgeReq struct {
	g.Meta   `path:"/judge" method:"POST" tags:"core" summary:"提交题目"`
	Code     string `json:"code" v:"required"`
	Username string `json:"username" v:"required"`
	Pid      string `json:"pid" v:"required"`
	Language string `json:"language" default:"cpp"`
	Cid      string `json:"cid"`
}

type JudgeRes struct {
	Sid    string `json:"sid"`
	Result string `json:"result"`
}
