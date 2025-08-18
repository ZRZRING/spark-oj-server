package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/submission"
)

func (c *ControllerSubmission) GetSubmission(ctx context.Context, req *submission.GetSubmissionReq) (res *submission.GetSubmissionRes, err error) {
	res = &submission.GetSubmissionRes{}
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
