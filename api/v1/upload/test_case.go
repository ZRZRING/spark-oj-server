package upload

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type TestCaseReq struct {
	g.Meta    `path:"/upload" method:"POST" tags:"core" mime:"multipart/form-data" summary:"提交文件"`
	TestCases []*ghttp.UploadFile `p:"TestCases" v:"required"`
	Pid       string              `p:"pid" v:"required"`
}

type TestCaseRes struct{}
