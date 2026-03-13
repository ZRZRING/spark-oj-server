package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProtectedReq struct {
	g.Meta `path:"/protected" method:"GET" tags:"admin" summary:"根据 JWT 令牌，获取当前登录的用户名和角色信息"`
}

type ProtectedRes struct {
	Username string `json:"username"`
	UserRole string `json:"user_role" dc:"用户角色,admin或user"`
}
