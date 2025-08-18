package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetContestList(ctx context.Context, req *contest.GetContestListReq) (res *contest.GetContestListRes, err error) {
	res = &contest.GetContestListRes{}
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
