package problem

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteTestcaseReq struct {
	g.Meta    `path:"/testcases/{problemId}/{name}" method:"DELETE" tags:"problem" summary:"删除测试用例（成对删除 .in 和 .out）"`
	ProblemId string `p:"problemId" in:"path" v:"required#题目 ID 不能为空"`
	Name      string `p:"name" in:"path" v:"required#测试用例名称不能为空"`
}

type DeleteTestcaseRes struct{}
