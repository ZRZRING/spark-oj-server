// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"spark-oj-server/api/admin/v1"
)

type IAdminV1 interface {
	Protected(ctx context.Context, req *v1.ProtectedReq) (res *v1.ProtectedRes, err error)
}
