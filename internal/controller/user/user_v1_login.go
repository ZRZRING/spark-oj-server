package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/middleware"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	md := dao.UserBase.Ctx(ctx)

	// 检查是否存在用户
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil || cnt == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest)
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 获取用户信息
	user := &entity.UserBase{}
	err = md.Where("username", req.Username).Scan(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 密码解析
	md5Password, err := gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 判断密码
	if user.Password != md5Password {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "密码错误")
		g.Log().Error(ctx, err, user)
		return nil, err
	}

	// 生成 token
	token, err := middleware.GenToken(req.Username, user.Role)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res.Token = token

	return res, nil
}
