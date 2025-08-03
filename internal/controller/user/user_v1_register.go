package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	res = &v1.RegisterRes{}
	md := dao.UserBase.Ctx(ctx)

	// 判断重名用户
	cnt, err := md.Where("username", req.Username).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if cnt > 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "存在重名用户")
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 密码加密
	req.Password, err = gmd5.Encrypt(req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	data := &do.UserBase{}
	err = gconv.Scan(req, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// nickname 默认与 username 相同
	if data.Nickname == nil {
		data.Nickname = data.Username
	}

	msg, err := md.Insert(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return res, nil
}
