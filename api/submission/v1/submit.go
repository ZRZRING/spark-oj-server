package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SubmitReq struct {
	g.Meta   `path:"/submission" method:"POST" tag:"submission" summary:"提交题目"`
	Pid      string `json:"pid"`
	Username string `json:"username"`
	Cid      string `json:"cid"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type SubmitRes struct {
	Result string `json:"result"`
}
