package core

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type JudgeReq struct {
	g.Meta   `path:"/judge" method:"POST" tags:"core" summary:"提交题目"`
	Pid      string         `json:"pid" v:"required"`
	Cid      string         `json:"cid"`
	Username string         `json:"username" v:"required"`
	Code     string         `json:"code" v:"required"`
	Language enums.Language `json:"language" v:"required" default:"cpp"`
}

type JudgeRes struct {
	Sid    string            `json:"sid"`
	Result enums.JudgeStatus `json:"result"`
}
