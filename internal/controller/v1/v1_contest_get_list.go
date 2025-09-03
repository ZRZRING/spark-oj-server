package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetList(ctx context.Context, req *contest.GetListReq) (res *contest.GetListRes, err error) {
	res = &contest.GetListRes{}
	md := dao.Contest.Ctx(ctx)

	e := make([]*entity.Contest, 0)
	tot := 0
	err = md.OrderAsc("start_time").Page(req.Page, req.Size).ScanAndCount(&e, &tot, false)
	if err != nil {
		return nil, err
	}

	err = gconv.Scan(e, &res.Contests)
	if err != nil {
		return nil, err
	}
	res.Total = tot

	return res, nil
}
