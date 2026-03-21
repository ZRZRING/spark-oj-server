package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProtectedReq struct {
	g.Meta `path:"/protected" method:"GET" tags:"admin" summary:"JWT 认证"`
}

type ProtectedRes struct {
	Username   string     `json:"username"`
	UserRole   string     `json:"user_role" dc:"用户角色,admin或user"`
	ExpireTime gtime.Time `json:"expire_time"`
}
