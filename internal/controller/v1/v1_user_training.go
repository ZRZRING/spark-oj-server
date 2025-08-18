package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/user"
)

func (c *ControllerUser) Training(ctx context.Context, req *user.TrainingReq) (res *user.TrainingRes, err error) {
	res = &user.TrainingRes{}
	md := dao.UserBase.Ctx(ctx)

	// 获取 URL 路径参数
	r := g.RequestFromCtx(ctx)
	username := gconv.String(r.Get("username").Val())

	// 从数据库获取用户训练数据
	user := &entity.UserBase{}
	err = md.Where("username", username).Scan(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 处理返回信息
	err = gconv.Scan(user, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 获取第三方训练信息
	if user.CfId != "" {
		data := service.GetUserInfo(ctx, user.CfId)
		res.CFRating = data.Rating
	}

	return res, nil
}
