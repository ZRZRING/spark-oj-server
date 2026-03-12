package problem

import "github.com/gogf/gf/v2/frame/g"

type GetPageReq struct {
	g.Meta `path:"/problems" method:"GET" summary:"获取题目列表" tags:"problem"`
	Page   int `p:"page" v:"required#页码不能为空"`
	Size   int `p:"size" v:"required#每页数量不能为空"`
}

type GetPageRes struct {
	Total    int `json:"total" dc:"题目总数"`
	Problems []*struct {
		Pid       string `json:"pid"`
		Title     string `json:"title"`
		JudgeType int    `json:"judge_type"`
		Rating    int    `json:"rating"`
	} `json:"problems" dc:"题目列表"`
}
