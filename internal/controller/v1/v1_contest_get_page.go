package v1

import (
	"context"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/contest"
)

func (c *ControllerContest) GetPage(ctx context.Context, req *contest.GetPageReq) (res *contest.GetPageRes, err error) {
	res = &contest.GetPageRes{}
	md := dao.Contest.Ctx(ctx)

	e := make([]*entity.Contest, req.Size)
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

	for i, item := range res.Contests {
		if e[i].Password != "" {
			item.Visibility = enums.ContestVisibilityPrivate
		} else {
			item.Visibility = enums.ContestVisibilityPublic
		}
	}

	return res, nil
}
