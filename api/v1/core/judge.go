package core

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type JudgeReq struct {
	g.Meta   `path:"/judge" method:"POST" tags:"core" summary:"提交题目"`
	Code     string         `json:"code" v:"required"`
	Username string         `json:"username" v:"required"`
	Pid      string         `json:"pid" v:"required"`
	Language enums.Language `json:"language" v:"enums" default:"cpp"`
	Cid      string         `json:"cid"`
}

type JudgeRes struct {
	Sid    string                  `json:"sid"`
	Result enums.JudgeResultStatus `json:"result" v:"enums"`
}
