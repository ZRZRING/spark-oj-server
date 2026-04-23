package submission

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageReq struct {
	g.Meta    `path:"/submissions" method:"GET" tags:"submission" summary:"分页获取提交列表"`
	Page      int    `p:"page" v:"required#页码不能为空"`
	Size      int    `p:"size" v:"required#每页数量不能为空"`
	Username  string `p:"username"`
	ProblemId string `p:"problemId"`
	ContestId string `p:"contestId"`
}

type GetPageRes struct {
	Total       int            `json:"total"`
	Submissions []*GetPageItem `json:"submissions"`
}

type GetPageItem struct {
	SubmissionId string `json:"submission_id"`
	ProblemId    string `json:"problem_id"`
	ContestId    string `json:"contest_id"`
	Username     string `json:"username"`
	Result       string `json:"result"`
	Language     string `json:"language"`
	MemoryCost   string `json:"memory_cost" dc:"MB"`
	TimeCost     string `json:"time_cost" dc:"ms"`
	CreateTime   string `json:"create_time"`
}
