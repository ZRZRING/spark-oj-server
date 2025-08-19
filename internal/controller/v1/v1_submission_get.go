package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/submission"
)

func (c *ControllerSubmission) Get(ctx context.Context, req *submission.GetReq) (res *submission.GetRes, err error) {
	res = &submission.GetRes{}
	md := dao.Submission.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	sid := gconv.String(r.Get("sid").Val())

	e := &entity.Submission{}
	err = md.Where("sid", sid).Scan(e)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(e, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
