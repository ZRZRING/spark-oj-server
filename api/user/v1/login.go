package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"user" summary:"用户登录"`
	Username string `json:"username" v:"required|length:3,16#请输入账号|账号长度为:{min}到:{max}位" dc:""`
	Password string `json:"password" v:"required|length:3,16#请输入密码|密码长度为:{min}到:{max}位" dc:""`
}

type LoginRes struct {
	Token string `json:"token" dc:""`
}
