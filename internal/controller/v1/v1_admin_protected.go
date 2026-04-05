package v1

import (
	"context"
	"spark-oj/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/admin"
)

func (c *ControllerAdmin) Protected(ctx context.Context, req *admin.ProtectedReq) (res *admin.ProtectedRes, err error) {
	res = &admin.ProtectedRes{
		Username:   gconv.String(ctx.Value(middleware.CtxUsername)),
		UserRole:   gconv.String(ctx.Value(middleware.CtxUserRole)),
		ExpireTime: *gconv.GTime(ctx.Value(middleware.CtxExpireTime)),
	}

	// 输出 JWT 有效时间
	g.Log().Info(ctx, "JWT 有效时间:", middleware.ExpireTime)

	return
}
