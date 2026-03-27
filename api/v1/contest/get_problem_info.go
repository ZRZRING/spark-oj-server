package contest

import (
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type GetProblemInfoReq struct {
	g.Meta `path:"/contest/{cid}/problem/{pid}" method:"GET" tags:"contest" summary:"获取比赛题目详情"`
	Cid    string `p:"cid" v:"required#比赛 ID 不能为空" dc:"比赛 ID"`
	Pid    string `p:"pid" v:"required#题目 ID 不能为空" dc:"题目 ID"`
}

type GetProblemInfoRes struct {
	Pid          string          `json:"pid"`
	Title        string          `json:"title"`
	JudgeType    enums.JudgeType `json:"judge_type"`
	TimeLimit    int             `json:"time_limit"`
	MemoryLimit  int             `json:"memory_limit"`
	Rating       int             `json:"rating"`
	Content      string          `json:"content"`
	Cid          string          `json:"cid"`
	ContestTitle string          `json:"contest_title"`
}