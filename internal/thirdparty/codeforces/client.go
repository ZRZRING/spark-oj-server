package codeforces

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"spoj/internal/consts"
)

func GetUserInfo(ctx g.Ctx, cfId string) *UserInfo {
	r, err := g.Client().Get(ctx, consts.CodeForcesAPI, g.Map{
		"handles": cfId,
	})
	defer func(r *gclient.Response) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	var response struct {
		Status string `json:"status"`
		Result g.List `json:"result"`
	}
	data := &UserInfo{}
	// g.Log().Debug(ctx, r.ReadAllString())
	err = gconv.Scan(r.ReadAllString(), &response)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	err = gconv.Scan(response.Result[0], &data)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	g.Log().Info(ctx, data)

	return data
}
