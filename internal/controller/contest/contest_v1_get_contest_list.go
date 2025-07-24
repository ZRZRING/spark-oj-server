package contest

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/contest/v1"
)

func (c *ControllerV1) GetContestList(ctx context.Context, req *v1.GetContestListReq) (res *v1.GetContestListRes, err error) {
	res = &v1.GetContestListRes{}
	md := dao.Problem.Ctx(ctx)
	var contests []*entity.Contest
	err = md.
		OrderAsc("start_time").
		Page(req.Page, req.Size).
		Scan(contests)
	err = gconv.Scan(contests)
	return res, nil
}
