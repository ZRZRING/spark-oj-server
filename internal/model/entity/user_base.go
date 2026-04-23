// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure for table user_base.
type UserBase struct {
	Username  string      `json:"username"  orm:"username"  ` //
	CreateAt  *gtime.Time `json:"create_at" orm:"create_at" ` //
	UpdateAt  *gtime.Time `json:"update_at" orm:"update_at" ` //
	DeleteAt  *gtime.Time `json:"delete_at" orm:"delete_at" ` //
	Password  string      `json:"password"  orm:"password"  ` //
	Nickname  string      `json:"nickname"  orm:"nickname"  ` //
	Role      string      `json:"role"      orm:"role"      ` //
	Solved    string      `json:"solved"    orm:"solved"    ` //
	Submitted string      `json:"submitted" orm:"submitted" ` //
	Rating    int         `json:"rating"    orm:"rating"    ` //
	Extra     string      `json:"extra"     orm:"extra"     ` //
}
