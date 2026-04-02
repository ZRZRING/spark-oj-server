package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetInfo(ctx context.Context, req *contest.GetInfoReq) (res *contest.GetInfoRes, err error) {
	res = &contest.GetInfoRes{}
	md := dao.Contest.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	contestId := gconv.String(r.Get("contestId").Val())

	e := &entity.Contest{}
	err = md.Where("contestId", contestId).Scan(e)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(e, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
