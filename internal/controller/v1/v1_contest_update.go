package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) Update(ctx context.Context, req *contest.UpdateReq) (res *contest.UpdateRes, err error) {
	res = &contest.UpdateRes{}
	md := dao.Contest.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	cid := gconv.String(r.Get("cid").Val())

	data := &do.Contest{}
	err = gconv.Struct(req, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	data.Cid = cid
	msg, err := md.OnConflict("cid").Save(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Info(ctx, msg)

	return res, nil
}
