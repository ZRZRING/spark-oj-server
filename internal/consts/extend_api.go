package consts

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	CodeForcesAPI = g.Cfg().MustGet(gctx.New(), "thirdParty.codeforces.user.baseUrl").String()
	JudgeAPI      = g.Cfg().MustGet(gctx.New(), "thirdParty.judge.baseUrl").String()
)
