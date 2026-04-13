package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"spark-oj/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	// 设置语言
	g.I18n().SetLanguage("zh-CN")

	// 启动服务
	cmd.Main.Run(gctx.GetInitCtx())
}
