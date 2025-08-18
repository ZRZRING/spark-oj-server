package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/user"
)

func (c *ControllerUser) Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error) {
	res = &user.ProfileRes{}
	md := dao.UserBase.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	username := gconv.String(r.Get("username").Val())

	user := &entity.UserBase{}
	err = md.Where("username", username).Scan(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Struct(user, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
