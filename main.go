package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"spark-oj-server/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func main() {
	// 设置语言
	g.I18n().SetLanguage("zh-CN")

	// 测试数据库连接
	err := g.DB().PingMaster()
	if err != nil {
		panic(err)
	}

	// 启动服务
	cmd.Main.Run(gctx.GetInitCtx())
}
