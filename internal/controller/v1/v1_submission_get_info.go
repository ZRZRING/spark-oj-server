package v1

import (
	"context"
	"fmt"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/submission"
)

func (c *ControllerSubmission) GetInfo(ctx context.Context, req *submission.GetInfoReq) (res *submission.GetInfoRes, err error) {
	res = &submission.GetInfoRes{}
	md := dao.Submission.Ctx(ctx)

	r := g.RequestFromCtx(ctx)
	sid := gconv.String(r.Get("sid").Val())

	e := &entity.Submission{}
	err = md.Where("sid", sid).Scan(e)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	res = &submission.GetInfoRes{
		Sid:        gconv.String(e.Sid),
		Pid:        gconv.String(e.Pid),
		Cid:        gconv.String(e.Cid),
		Username:   e.Username,
		Result:     e.Result,
		Language:   e.Language,
		MemoryCost: fmt.Sprintf("%.2f", float64(e.MemoryCost)/1024.0/1024.0),
		TimeCost:   fmt.Sprintf("%d", e.TimeCost),
		Code:       e.Code,
		CreateAt:   e.CreateAt.Format("2006-01-02 15:04:05"),
	}
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return res, nil
}
