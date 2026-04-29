package cmd

import (
	"context"
	"spark-oj/internal/router"
	"spark-oj/pkg/sshtunnel"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			if err := sshtunnel.Setup(); err != nil {
				g.Log().Fatal(ctx, "ssh tunnel setup failed:", err)
			}
			timezoneConfig, err := g.Cfg().Get(ctx, "server.timezone")
			if err != nil {
				g.Log().Info(ctx, "server time zone failed:", err)
			}
			timezone := timezoneConfig.String()
			if timezone == "" {
				timezone = "Asia/Shanghai"
			}
			if err := gtime.SetTimeZone(timezone); err != nil {
				g.Log().Fatal(ctx, "set time zone failed:", err)
			}
			return router.Bind(ctx, parser)
		},
	}
)
