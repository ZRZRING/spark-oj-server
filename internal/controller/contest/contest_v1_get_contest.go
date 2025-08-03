package contest

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/contest/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
)

func (c *ControllerV1) GetContest(ctx context.Context, req *v1.GetContestReq) (res *v1.GetContestRes, err error) {
	res = &v1.GetContestRes{}
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
