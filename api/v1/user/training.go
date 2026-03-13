package user

import "github.com/gogf/gf/v2/frame/g"

type TrainingReq struct {
	g.Meta `path:"/training/{username}" method:"GET" tags:"user" summary:"获取用户训练数据"`
}
type TrainingRes struct {
	Submitted int `json:"submitted"`
	Rating    int `json:"rating"`
	Solved    int `json:"solved"`
	CFRating  int `json:"cf_rating"`
	ATCRating int `json:"atc_rating"`
}
