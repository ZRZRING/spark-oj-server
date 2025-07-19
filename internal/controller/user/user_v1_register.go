package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/consts"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	g.Log().Info(ctx, req)
	md := dao.UserBase.Ctx(ctx)
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		err = consts.ErrUserExist
		return nil, err
	}
	g.Log().Info(ctx, cnt)
	req.Password, err = gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data := &do.UserBase{}
	err = gconv.Struct(req, data)
	if err != nil {
		return nil, err
	}
	if data.Nickname == nil {
		data.Nickname = data.Username
	}
	msg, err := md.Insert(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)
	return
}
