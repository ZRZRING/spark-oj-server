package contest

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spoj/api/contest/v1"
	"spoj/internal/dao"
	"spoj/internal/model/entity"
)

func (c *ControllerV1) GetContest(ctx context.Context, req *v1.GetContestReq) (res *v1.GetContestRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := dao.Contest.Ctx(ctx)
	contest := &entity.Contest{}
	cid := gconv.String(r.Get("cid").Val())

	err = md.Where("cid", cid).Scan(&contest)
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
