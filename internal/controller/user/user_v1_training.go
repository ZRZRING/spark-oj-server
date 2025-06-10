package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"spoj/api/user/v1"
	"spoj/internal/consts"
	"spoj/internal/dao"
	"spoj/internal/model/entity"
)

func (c *ControllerV1) Training(ctx context.Context, req *v1.TrainingReq) (res *v1.TrainingRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := dao.UserBase.Ctx(ctx)
	user := &entity.UserBase{}
	username := gconv.String(r.Get("username").Val())

	err = md.Where("username", username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(user, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	if user.CfId != "" {
		r, err := g.Client().Get(ctx, consts.CodeForcesAPI, g.Map{
			"handles": user.CfId,
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
		var data struct {
			LastOnlineTimeSeconds gtime.Time `json:"lastOnlineTimeSeconds"`
			Rating                int        `json:"rating"`
			MaxRating             int        `json:"maxRating"`
		}
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
		res.CFRating = data.Rating
	}

	return res, nil
}
