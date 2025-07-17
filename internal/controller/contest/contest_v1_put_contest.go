package contest

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark-oj-server/api/contest/v1"
)

func (c *ControllerV1) PutContest(ctx context.Context, req *v1.PutContestReq) (res *v1.PutContestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
