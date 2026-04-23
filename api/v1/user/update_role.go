package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateRoleReq struct {
	g.Meta   `path:"/users/role/{username}" method:"PUT" tags:"user" summary:"更新用户权限"`
	Username string `p:"username" in:"path" v:"required#用户名不能为空"`
	Role     string `p:"role" v:"required|in:user,admin,locked,root#角色不能为空|角色不存在"`
}

type UpdateRoleRes struct {
}
