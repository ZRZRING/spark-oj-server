package contest

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/contest/v1"
)

func (c *ControllerV1) GetContestList(ctx context.Context, req *v1.GetContestListReq) (res *v1.GetContestListRes, err error) {
	res = &v1.GetContestListRes{}
	md := dao.Problem.Ctx(ctx)

	contests := make([]*entity.Contest, 0)
	tot := new(int)
	err = md.OrderAsc("start_time").Page(req.Page, req.Size).ScanAndCount(contests, tot, false)
	if err != nil {
		return nil, err
	}

	err = gconv.Scan(contests, res)
	if err != nil {
		return nil, err
	}
	res.Total = *tot

	return res, nil
}
