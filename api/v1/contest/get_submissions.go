package contest

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSubmissionsReq struct {
	g.Meta    `path:"/contest/{contestId}/submissions" method:"GET" tags:"contest" summary:"获取比赛提交列表"`
	Page      int    `p:"page" v:"required#页码不能为空"`
	Size      int    `p:"size" v:"required#每页数量不能为空"`
	ContestId string `p:"contestId"`
}

type GetSubmissionsRes struct {
	Total       int                   `json:"total"`
	Submissions []*GetSubmissionsItem `json:"submissions"`
}

type GetSubmissionsItem struct {
	SubmissionId string `json:"submissionId"`
	ProblemId    string `json:"problemId"`
	ContestId    string `json:"contestId"`
	Username     string `json:"username"`
	Result       string `json:"result"`
	Language     string `json:"language"`
	MemoryCost   string `json:"memory_cost" dc:"MB"`
	TimeCost     string `json:"time_cost" dc:"ms"`
	CreateTime   string `json:"create_time"`
}
