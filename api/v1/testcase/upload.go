package testcase

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta    `path:"/upload/testcases" method:"POST" tags:"core" summary:"提交测试用例"`
	ProblemId string            `p:"problemId" v:"required"`
	Testcases ghttp.UploadFiles `p:"testcases" type:"file" v:"required"`
}

type UploadRes struct {
	Path string `json:"path"`
}
