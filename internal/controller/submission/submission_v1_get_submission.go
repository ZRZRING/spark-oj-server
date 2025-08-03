package submission

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/submission/v1"
)

func (c *ControllerV1) GetSubmission(ctx context.Context, req *v1.GetSubmissionReq) (res *v1.GetSubmissionRes, err error) {
	res = &v1.GetSubmissionRes{}
	md := dao.Submission.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	sid := gconv.String(r.Get("sid").Val())

	data := &entity.Submission{}
	err = md.Where("sid", sid).Scan(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	err = gconv.Scan(data, res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
