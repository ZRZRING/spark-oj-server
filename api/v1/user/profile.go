package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta   `path:"/profile/{username}" method:"GET" tags:"user" summary:"获取用户信息"`
	Username string `p:"username" in:"path" v:"required#用户名不能为空"`
}

type ProfileRes struct {
	Nickname   string `json:"nickname"`
	Company    string `json:"company"`
	Department string `json:"department"`
	Major      string `json:"major"`
	Class      string `json:"class"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	Avatar     string `json:"avatar"`
}
