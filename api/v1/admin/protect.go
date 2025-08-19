package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProtectedReq struct {
	g.Meta `path:"/protected" method:"GET" tags:"admin" summary:"测试 JWT 中间件"`
}

type ProtectedRes struct {
	Username string `json:"username" dc:""`
	UserRole string `json:"user_role" dc:""`
}
