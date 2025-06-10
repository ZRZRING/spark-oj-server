package v1

import "github.com/gogf/gf/v2/frame/g"

type TrainingReq struct {
	g.Meta `path:"/training/{username}" method:"get" tags:"user" summary:"获取用户训练数据"`
}
type TrainingRes struct {
	Submitted int `json:"submitted" dc:""`
	Rating    int `json:"rating" dc:""`
	Solved    int `json:"solved" dc:""`
	CFRating  int `json:"cf_rating" dc:""`
	ATCRating int `json:"atc_rating" dc:""`
}
