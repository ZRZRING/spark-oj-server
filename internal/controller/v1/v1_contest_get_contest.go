package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetContest(ctx context.Context, req *contest.GetContestReq) (res *contest.GetContestRes, err error) {
	res = &contest.GetContestRes{}
	md := dao.Contest.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	cid := gconv.String(r.Get("cid").Val())

	contest := &entity.Contest{}
	err = md.Where("cid", cid).Scan(contest)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(contest, &res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
