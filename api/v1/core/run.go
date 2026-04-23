package core

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RunReq struct {
	g.Meta   `path:"/run" method:"POST" tags:"core" summary:"运行代码"`
	Code     string `json:"code"`
	Input    string `json:"input"`
	Language string `json:"language"`
}

type RunRes struct {
	Status     string `json:"status"`
	Output     string `json:"output"`
	Error      string `json:"error"`
	Time       int64  `json:"time"`
	Memory     int64  `json:"memory"`
	ExitStatus int64  `json:"exit_status"`
}
