package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/contest"
)

func (c *ControllerContest) GetInfo(ctx context.Context, req *contest.GetInfoReq) (res *contest.GetInfoRes, err error) {
	res = &contest.GetInfoRes{}
	md := dao.Contest.Ctx(ctx)

	e := &entity.Contest{}
	err = md.Where("contest_id", req.ContestId).Scan(e)
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
