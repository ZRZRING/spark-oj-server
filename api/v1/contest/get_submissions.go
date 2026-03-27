package contest

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSubmissionsReq struct {
	g.Meta `path:"/contest/{cid}/submissions" method:"GET" tags:"contest" summary:"获取比赛提交列表"`
	Page   int    `p:"page" v:"required#页码不能为空"`
	Size   int    `p:"size" v:"required#每页数量不能为空"`
	Cid    string `p:"cid"`
}

type GetSubmissionsRes struct {
	Total       int                   `json:"total"`
	Submissions []*GetSubmissionsItem `json:"submissions"`
}

type GetSubmissionsItem struct {
	Sid        string `json:"sid"`
	Pid        string `json:"pid"`
	Cid        string `json:"cid"`
	Username   string `json:"username"`
	Result     string `json:"result"`
	Language   string `json:"language"`
	MemoryCost string `json:"memory_cost" dc:"MB"`
	TimeCost   string `json:"time_cost" dc:"ms"`
	CreateTime string `json:"create_time"`
}
