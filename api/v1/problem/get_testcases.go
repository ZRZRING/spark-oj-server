package problem

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetTestcasesReq struct {
	g.Meta    `path:"/testcases/{problemId}" method:"GET" tags:"problem" summary:"查询测试用例"`
	ProblemId string `p:"problemId" in:"path" v:"required#题目 ID 不能为空"`
}

type GetTestcasesRes struct {
	Testcases []TestcaseItem `json:"testcases"`
}

type TestcaseItem struct {
	Name      string `json:"name"`
	InputSize int64  `json:"inputSize"`
	OutputSize int64 `json:"outputSize"`
}
