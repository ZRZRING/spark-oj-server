package contest

import (
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
)

type GetProblemInfoReq struct {
	g.Meta    `path:"/contest/{contestId}/problem/{problemId}" method:"GET" tags:"contest" summary:"获取比赛题目详情"`
	ContestId string `p:"contestId" in:"path" v:"required#比赛 ID 不能为空"`
	ProblemId string `p:"problemId" in:"path" v:"required#题目 ID 不能为空"`
}

type GetProblemInfoRes struct {
	ProblemId    string          `json:"problem_id"`
	Title        string          `json:"title"`
	JudgeType    enums.JudgeType `json:"judge_type"`
	TimeLimit    int             `json:"time_limit"`
	MemoryLimit  int             `json:"memory_limit"`
	Rating       int             `json:"rating"`
	Content      string          `json:"content"`
	ContestId    string          `json:"contest_id"`
	ContestTitle string          `json:"contest_title"`
}
