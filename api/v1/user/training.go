package user

import "github.com/gogf/gf/v2/frame/g"

type TrainingReq struct {
	g.Meta   `path:"/training/{username}" method:"GET" tags:"user" summary:"获取用户训练数据"`
	Username string `p:"username" in:"path" v:"required#用户名不能为空"`
}
type TrainingRes struct {
	Submitted int `json:"submitted"`
	Solved    int `json:"solved"`
	Rating    int `json:"rating"`
	// CFRating  int `json:"cf_rating"`
	// ATCRating int `json:"atc_rating"`
}
