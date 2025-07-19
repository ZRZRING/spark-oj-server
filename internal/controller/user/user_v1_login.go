package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/consts"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/middleware"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	md := dao.UserBase.Ctx(ctx)
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if cnt == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "没有对应用户")
		g.Log().Error(ctx, err)
		return nil, err
	}
	user := &entity.UserBase{}
	err = md.Where("username", req.Username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	md5Password, err := gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if user.Password != md5Password {
		err = consts.ErrInvalidPassword
		g.Log().Error(ctx, err, user)
		return nil, err
	}
	token, err := middleware.GenToken(req.Username, user.Role)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res = &v1.LoginRes{
		Token: token,
	}
	return res, nil
}
