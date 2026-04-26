package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/contest"
)

func (c *ControllerContest) Update(ctx context.Context, req *contest.UpdateReq) (res *contest.UpdateRes, err error) {
	res = &contest.UpdateRes{}
	md := dao.Contest.Ctx(ctx)

	d := &do.Contest{}
	err = gconv.Struct(req, d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	d.ContestId = req.ContestId
	msg, err := md.OnConflict("contest_id").Save(d)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return res, nil
}
