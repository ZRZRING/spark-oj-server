package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSubmissionListReq struct {
	g.Meta `path:"/submissions" method:"GET" tags:"submission" summary:"分页获取提交列表"`
	Page   int `json:"page" dc:""`
	Size   int `json:"size" dc:""`
}

type GetSubmissionListRes struct {
	Total      int `json:"total" dc:""`
	Submission []*struct {
		Sid        string `json:"sid"`
		Pid        string `json:"pid"`
		Username   string `json:"username"`
		Cid        string `json:"cid"`
		Result     string `json:"result"`
		Language   string `json:"language"`
		MemoryCost int    `json:"memory_cost"`
		TimeCost   int    `json:"time_cost"`
		CreateTime int    `json:"create_time"`
	} `json:"submission" dc:""`
}
