package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/middleware"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"spark-oj-server/api/v1/user"
)

func (c *ControllerUser) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	res = &user.LoginRes{}
	md := dao.UserBase.Ctx(ctx)

	// 检查是否存在用户
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil || cnt == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest)
		g.Log().Error(ctx, err)
		return nil, err
	}

	e := &entity.UserBase{}
	err = md.Where("username", req.Username).Scan(e)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 获取 md5 加密后的信息
	md5Password, err := gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 判断密码
	if e.Password != md5Password {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "密码错误")
		g.Log().Error(ctx, err, e)
		return nil, err
	}

	// 生成 token
	token, err := middleware.GenToken(req.Username, e.Role)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	res.Token = token

	return res, nil
}
