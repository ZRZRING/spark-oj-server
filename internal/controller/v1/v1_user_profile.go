package v1

import (
	"context"
	"spark-oj/api/v1/user"
)

func (c *ControllerUser) Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error) {
	res = &user.ProfileRes{}

	return res, nil
}
