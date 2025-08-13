package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/consts"
)

type UserInfo struct {
	LastOnlineTimeSeconds gtime.Time `json:"lastOnlineTimeSeconds"`
	Rating                int        `json:"rating"`
	MaxRating             int        `json:"maxRating"`
}

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
	// g.Log().Debug(ctx, r.ReadAllString())
	if err != nil {
		g.Log().Error(ctx, err)
	}

	var response struct {
		Status string `json:"status"`
		Result g.List `json:"result"`
	}
	err = gconv.Scan(r.ReadAllString(), &response)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	
	data := &UserInfo{}
	err = gconv.Scan(response.Result[0], &data)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	g.Log().Info(ctx, data)

	return data
}
