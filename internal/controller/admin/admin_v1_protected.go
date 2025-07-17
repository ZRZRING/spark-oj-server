package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/middleware"

	"spark-oj-server/api/admin/v1"
)

func (c *ControllerV1) Protected(ctx context.Context, req *v1.ProtectedReq) (res *v1.ProtectedRes, err error) {
	username := gconv.String(ctx.Value(middleware.CtxUsername))
	userRole := gconv.String(ctx.Value(middleware.CtxUserRole))
	res = &v1.ProtectedRes{
		Username: username,
		UserRole: userRole,
	}
	return
}
