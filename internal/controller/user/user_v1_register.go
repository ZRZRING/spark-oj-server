package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"spoj/api/user/v1"
	"spoj/internal/consts"
	"spoj/internal/dao"
	"spoj/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	g.Log().Info(ctx, req)

	// 校验参数
	err = g.Validator().Data(req).Run(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	md := dao.UserBase.Ctx(ctx)

	// 检查用户是否存在
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		err = consts.ErrUserExist
		return nil, err
	}
	g.Log().Info(ctx, cnt)

	// 往数据库添加用户
	md5Password, err := gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	msg, err := md.Insert(do.UserBase{
		Username: req.Username,
		Password: md5Password,
		Nickname: req.Username,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return
}
