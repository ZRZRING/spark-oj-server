package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{}
	md := dao.UserBase.Ctx(ctx)

	// 获取请求参数
	r := g.RequestFromCtx(ctx)
	username := gconv.String(r.Get("username").Val())

	// 获取用户信息
	user := &entity.UserBase{}
	err = md.Where("username", username).Scan(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 绑定返回信息
	err = gconv.Struct(user, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
