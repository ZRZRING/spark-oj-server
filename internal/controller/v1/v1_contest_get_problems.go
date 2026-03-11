package v1

import (
	"context"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetProblems(ctx context.Context, req *contest.GetProblemsReq) (res *contest.GetProblemsRes, err error) {
	res = &contest.GetProblemsRes{}
	md := dao.Contest.Ctx(ctx)

	problemIds := make([]int, 0)
	err = md.Where("cid", req.Cid).Fields("problems").Scan(&problemIds)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	problems := make([]contest.GetProblemsItem, 0)
}
