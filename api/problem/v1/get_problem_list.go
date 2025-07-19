package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProblemListReq struct {
	g.Meta `path:"/problems" method:"get" summary:"获取题目列表" tags:"problem"`
	Page   int `json:"page" v:"required#页码不能为空"`
	Size   int `json:"size" v:"required#每页数量不能为空"`
}

type GetProblemListRes struct {
	Problems []*Problem `json:"problems" dc:"题目列表"`
}

type Problem struct {
	Pid   string `json:"pid"`
	Title string `json:"title"`
}
