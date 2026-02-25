package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) Ranking(ctx context.Context, req *contest.RankingReq) (res *contest.RankingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
