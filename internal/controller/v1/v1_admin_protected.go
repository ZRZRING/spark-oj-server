package v1

import (
	"context"
	"spark-oj-server/internal/middleware"

	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/admin"
)

func (c *ControllerAdmin) Protected(ctx context.Context, req *admin.ProtectedReq) (res *admin.ProtectedRes, err error) {
	res = &admin.ProtectedRes{}

	username := gconv.String(ctx.Value(middleware.CtxUsername))
	userRole := gconv.String(ctx.Value(middleware.CtxUserRole))

	res.Username = username
	res.UserRole = userRole

	return
}
