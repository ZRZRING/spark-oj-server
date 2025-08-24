package router

import (
	"context"
	v1 "spark-oj-server/internal/controller/v1"
	"spark-oj-server/internal/middleware"

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
			v1.NewUser(),
			v1.NewProblem(),
			v1.NewContest(),
			v1.NewCore(),
			v1.NewSubmission(),
			v1.NewUpload(),
		)
	})

	// Protected endpoints requiring JWT authentication
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Middleware(middleware.JWTAuth)
		group.Middleware(middleware.PermissionAuth)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			v1.NewAdmin(),
		)
	})

	s.SetServerRoot("/resource/public")
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	s.SetPort(8000)

	s.Run()

	return
}
