package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta     `path:"/register" method:"post" tags:"user" summary:"用户注册"`
	Username   string `json:"username" v:"required|length:3,16#请输入账号|账号长度为:{min}到:{max}位" dc:""`
	Password   string `json:"password" v:"required|length:3,16#请输入密码|密码长度为:{min}到:{max}位" dc:""`
	RePassword string `json:"re_password" v:"required|length:3,16|same:Password#请再次输入密码|密码长度为:{min}到:{max}位|与第一次输入密码不相符" dc:""`
}

type RegisterRes struct{}
