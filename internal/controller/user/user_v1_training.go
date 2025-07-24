package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/user/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/internal/thirdparty/codeforces"
)

func (c *ControllerV1) Training(ctx context.Context, req *v1.TrainingReq) (res *v1.TrainingRes, err error) {
	res = &v1.TrainingRes{}
	md := dao.UserBase.Ctx(ctx)

	// 获取 URL 路径参数
	r := g.RequestFromCtx(ctx)
	user := &entity.UserBase{}
	username := gconv.String(r.Get("username").Val())

	// 从数据库获取用户训练数据
	err = md.Where("username", username).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 绑定信息
	err = gconv.Scan(user, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 获取第三方训练信息
	if user.CfId != "" {
		data := codeforces.GetUserInfo(ctx, user.CfId)
		res.CFRating = data.Rating
	}

	return res, nil
}
