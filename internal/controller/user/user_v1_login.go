package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"spoj/api/user/v1"
	"spoj/internal/consts"
	"spoj/internal/dao"
	"spoj/internal/middleware"
	"spoj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 校验请求
	err = g.Validator().Data(req).Run(ctx)
	if err != nil {
		g.Log().Error(ctx, "登录信息校验失败", err)
		return
	}

	md := dao.UserBase.Ctx(ctx)

	// 检查用户是否存在
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil {
		g.Log().Error(ctx, "<UNK>", err)
		return
	}
	if cnt == 0 {
		err = consts.ErrUserNotExist
		g.Log().Error(ctx, err)
		return
	}

	// 绑定数据库内用户信息
	user := &entity.UserBase{}
	err = md.Where("username", req.Username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 检查密码是否正确
	md5Password, err := gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if user.Password != md5Password {
		err = consts.ErrInvalidPassword
		g.Log().Error(ctx, err, user)
		return
	}

	// 生成 token
	token, err := middleware.GenToken(req.Username, user.Role)
	if err != nil {
		g.Log().Error(ctx, "生成 Token 失败", err)
		return
	}
	res = &v1.LoginRes{
		Token: token,
	}

	return
}
