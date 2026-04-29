package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{
		"localhost:5173",
		"localhost:5050",
		"127.0.0.1:5173",
		"127.0.0.1:5050",
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
