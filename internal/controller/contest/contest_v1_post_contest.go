package contest

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/api/contest/v1"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/do"
)

func (c *ControllerV1) PostContest(ctx context.Context, req *v1.PostContestReq) (res *v1.PostContestRes, err error) {
	res = &v1.PostContestRes{}
	md := dao.Contest.Ctx(ctx)

	data := &do.Contest{}
	err = gconv.Struct(req, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	msg, err := md.Insert(data)
	g.Log().Info(ctx, msg)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
