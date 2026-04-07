package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateUserRoleReq struct {
	g.Meta   `path:"/users/role" method:"PUT" tags:"admin" summary:"更新用户权限"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Role     string `p:"role" v:"required|in:user,admin#角色不能为空|角色必须是user或admin"`
}

type UpdateUserRoleRes struct {
}
