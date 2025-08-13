package user

import (
	"context"

	"spark-oj-server/api/user/v1"
)

func (c *ControllerV1) Submit(ctx context.Context, req *v1.SubmitReq) (res *v1.SubmitRes, err error) {
	res = &v1.SubmitRes{}
	
	return res, nil
}
