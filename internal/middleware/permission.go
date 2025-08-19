package middleware

import (
	"spark-oj-server/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func PermissionAuth(r *ghttp.Request) {
	username := r.GetCtx().Value(CtxUsername)
	if username == nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "用户名为空"))
	}

	md := dao.UserBase.Ctx(r.Context())
	cnt, err := md.Count("username", username)
	if err != nil || cnt == 0 {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "不存在该用户"))
	}

	r.Middleware.Next()
}
