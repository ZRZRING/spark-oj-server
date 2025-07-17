package contest

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/contest/v1"
)

func (c *ControllerV1) PostContest(ctx context.Context, req *v1.PostContestReq) (res *v1.PostContestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
