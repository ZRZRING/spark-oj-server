package core

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RunReq struct {
	g.Meta `path:"/run" method:"POST" tags:"core" mime:"application/json" tags:"core"`
	Code   string `json:"code"`
	StdIn  string `json:"std_in"`
}

type RunRes struct {
	Status string `json:"status"`
	StdOut string `json:"std_out"`
}
