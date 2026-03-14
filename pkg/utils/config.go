package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func GetConfigString(pattern string) string {
	return g.Cfg().MustGet(gctx.New(), pattern).String()
}
