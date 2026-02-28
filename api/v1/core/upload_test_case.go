package core

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadTestCaseReq struct {
	g.Meta    `path:"/upload/testcases" method:"POST" tags:"core" summary:"提交文件"`
	TestCases ghttp.UploadFiles `p:"test_cases" type:"file" v:"required"`
	Pid       int               `p:"pid" v:"required"`
}

type UploadTestCaseRes struct {
	Path string `json:"path"`
}
