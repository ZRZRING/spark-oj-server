package v1

import (
	"context"
	"fmt"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/submission"
)

func (c *ControllerSubmission) GetInfo(ctx context.Context, req *submission.GetInfoReq) (res *submission.GetInfoRes, err error) {
	res = &submission.GetInfoRes{}

	r := g.RequestFromCtx(ctx)
	submissionId := r.Get("submissionId").Int()

	e := &entity.Submission{}
	err = dao.Submission.Ctx(ctx).Where("submission_id", submissionId).Scan(e)
	if err != nil {
		return nil, err
	}

	var memoryCost, timeCost string
	if e.MemoryCost > 0 {
		memoryCost = fmt.Sprintf("%.2f", float64(e.MemoryCost)/1024.0/1024.0)
	}
	if e.TimeCost > 0 {
		timeCost = fmt.Sprintf("%d", e.TimeCost)
	}
	res = &submission.GetInfoRes{
		SubmissionId: gconv.String(e.SubmissionId),
		ProblemId:    gconv.String(e.ProblemId),
		ContestId:    gconv.String(e.ContestId),
		Username:     e.Username,
		Result:       e.Result,
		Language:     e.Language,
		MemoryCost:   memoryCost,
		TimeCost:     timeCost,
		Code:         e.Code,
		CreateTime:   e.CreateAt.Layout("2006-01-02 15:04:05"),
	}

	return res, nil
}
