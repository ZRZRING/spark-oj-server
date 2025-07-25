package cmd

import (
	"spark-oj-server/internal/router"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  router.Bind,
	}
)
