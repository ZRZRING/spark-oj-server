package problem

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadTestcaseReq struct {
	g.Meta    `path:"/upload/testcases" method:"POST" tags:"problem" summary:"提交测试用例"`
	ProblemId string            `p:"problemId" v:"required"`
	Testcases ghttp.UploadFiles `p:"testcases" type:"file" v:"required"`
}

type UploadTestcaseRes struct {
	Path string `json:"path"`
}
