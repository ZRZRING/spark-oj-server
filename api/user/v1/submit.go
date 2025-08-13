package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SubmitReq struct {
	g.Meta   `path:"/submission" method:"POST" tag:"submission" summary:"提交题目"`
	Pid      string `json:"pid" v:"required#题目编号不能为空"`
	Username string `json:"username" v:"required#标题不能为空"`
	Cid      string `json:"cid" v:"required#标题不能为空"`
	Language string `json:"language" v:"required#标题不能为空"`
	Code     string `json:"code" v:"required#标题不能为空"`
}

type SubmitRes struct {
	Result  string `json:"result"`
	FilesId g.Map  `json:"files_id"`
}
