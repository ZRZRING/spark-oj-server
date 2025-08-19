package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) Create(ctx context.Context, req *contest.CreateReq) (res *contest.CreateRes, err error) {
	res = &contest.CreateRes{}
	md := dao.Contest.Ctx(ctx)

	d := &do.Contest{}
	err = gconv.Struct(req, d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	msg, err := md.Insert(d)
	g.Log().Info(ctx, msg)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
