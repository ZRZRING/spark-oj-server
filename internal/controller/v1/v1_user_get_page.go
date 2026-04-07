package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"

	"spark-oj/api/v1/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerUser) GetPage(ctx context.Context, req *user.GetPageReq) (res *user.GetPageRes, err error) {
	res = &user.GetPageRes{}
	data := make([]*entity.UserBase, req.Size)
	tot := new(int)
	err = dao.UserBase.Ctx(ctx).
		Page(req.Page, req.Size).
		ScanAndCount(&data, tot, false)
	if err != nil {
		g.Log().Error(ctx, "数据库错误：", err)
		return nil, err
	}
	res.Total = *tot
	for _, item := range data {
		res.Users = append(res.Users, user.GetPageItem{
			Username:   item.Username,
			UserRole:   item.Role,
			CreateTime: item.CreateAt.String(),
			Rating:     gconv.String(item.Rating),
		})
	}
	return res, nil
}
