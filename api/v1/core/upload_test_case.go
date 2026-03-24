package core

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadTestCaseReq struct {
	g.Meta    `path:"/upload/testcases/{pid}" method:"POST" tags:"core" summary:"提交测试用例"`
	TestCases ghttp.UploadFiles `p:"test_cases" type:"file" v:"required"`
	Pid       string            `p:"problem_id" v:"required"`
}

type UploadTestCaseRes struct {
	Path string `json:"path"`
}
