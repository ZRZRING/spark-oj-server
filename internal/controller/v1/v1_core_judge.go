package v1

import (
	"context"

	"spark-oj-server/api/v1/core"
)

func (c *ControllerCore) Judge(ctx context.Context, req *core.JudgeReq) (res *core.JudgeRes, err error) {
	res = &core.JudgeRes{}

	return res, nil
}
