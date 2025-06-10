package router

import (
	"context"
	"spoj/internal/controller/admin"
	"spoj/internal/controller/problem"
	"spoj/internal/controller/user"
	"spoj/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func Bind(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			user.NewV1(),
			problem.NewV1(),
		)
	})

	// Protected endpoints requiring JWT authentication
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Middleware(middleware.JWTAuth)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			admin.NewV1(),
		)
	})

	s.SetServerRoot("/resource/public")
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	s.SetPort(8000)

	s.Run()

	return
}
