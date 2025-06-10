package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/profile/{username}" method:"get" tags:"user" summary:"获取用户信息"`
}

type ProfileRes struct {
	Nickname   string `json:"nickname" dc:""`
	Company    string `json:"company" dc:""`
	Department string `json:"department" dc:""`
	Major      string `json:"major" dc:""`
	Class      string `json:"class" dc:""`
	Email      string `json:"email" dc:""`
	Tel        string `json:"tel" dc:""`
	Avatar     string `json:"avatar" dc:""`
}
