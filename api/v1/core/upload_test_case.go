package core

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadTestCaseReq struct {
	g.Meta    `path:"/upload" method:"POST" tags:"upload" mime:"multipart/form-data" summary:"提交文件"`
	TestCases *ghttp.UploadFiles `p:"test_cases" v:"required"`
	Pid       int                `p:"pid" v:"required"`
}

type UploadTestCaseRes struct{}
