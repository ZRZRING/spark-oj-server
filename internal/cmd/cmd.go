package cmd

import (
	"context"

	"spark-oj/internal/router"
	"spark-oj/pkg/sshtunnel"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			if err := sshtunnel.Setup(); err != nil {
				g.Log().Fatal(ctx, "ssh tunnel setup failed:", err)
				return err
			}
			return router.Bind(ctx, parser)
		},
	}
)
