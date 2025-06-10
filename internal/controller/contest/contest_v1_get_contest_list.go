package contest

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spoj/api/contest/v1"
)

func (c *ControllerV1) GetContestList(ctx context.Context, req *v1.GetContestListReq) (res *v1.GetContestListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
