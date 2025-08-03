package contest

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"

	"spark-oj-server/api/contest/v1"
)

func (c *ControllerV1) PutContest(ctx context.Context, req *v1.PutContestReq) (res *v1.PutContestRes, err error) {
	res = &v1.PutContestRes{}
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
