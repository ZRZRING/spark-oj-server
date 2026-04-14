package v1

import (
	"context"
	"spark-oj/api/v1/user"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerUser) UpdateRole(ctx context.Context, req *user.UpdateRoleReq) (res *user.UpdateRoleRes, err error) {
	res = &user.UpdateRoleRes{}

	// 检查用户是否存在
	count, err := dao.UserBase.Ctx(ctx).Where("username", req.Username).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if count == 0 {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "用户不存在")
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 更新用户角色
	_, err = dao.UserBase.Ctx(ctx).
		Where("username", req.Username).
		Update(&do.UserBase{
			Role: req.Role,
		})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	g.Log().Infof(ctx, "用户 %s 的权限已更新为 %s", req.Username, req.Role)

	return res, nil
}
