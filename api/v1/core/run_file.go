package core

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type RunFileReq struct {
	g.Meta `path:"/run/file" method:"POST" tags:"core" mime:"multipart/form-data" tags:"core"`
	Code   *ghttp.UploadFile `p:"code"`
	StdIn  string            `p:"std_in"`
}

type RunFileRes struct {
	Status string `json:"status"`
	StdOut string `json:"std_out"`
}
