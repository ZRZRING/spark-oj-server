package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spoj/api/user/v1"
	"spoj/internal/dao"
	"spoj/internal/model/entity"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	md := dao.UserBase.Ctx(ctx)
	user := &entity.UserBase{}
	r := g.RequestFromCtx(ctx)
	username := gconv.String(r.Get("username").Val())
	err = md.Where("username", username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	err = gconv.Struct(user, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return res, nil
}
