// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Rid        string      `json:"rid"        orm:"rid"        ` //
	Permission string      `json:"permission" orm:"permission" ` //
	CreateAt   *gtime.Time `json:"create_at"  orm:"create_at"  ` //
	UpdateAt   *gtime.Time `json:"update_at"  orm:"update_at"  ` //
	DeleteAt   *gtime.Time `json:"delete_at"  orm:"delete_at"  ` //
}
