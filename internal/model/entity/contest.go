// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Contest is the golang structure for table contest.
type Contest struct {
	Cid         string      `json:"cid"         orm:"cid"         ` //
	Title       string      `json:"title"       orm:"title"       ` //
	Password    string      `json:"password"    orm:"password"    ` //
	Problems    string      `json:"problems"    orm:"problems"    ` //
	Description string      `json:"description" orm:"description" ` //
	StartTime   *gtime.Time `json:"start_time"  orm:"start_time"  ` //
	EndTime     *gtime.Time `json:"end_time"    orm:"end_time"    ` //
	LockTime    *gtime.Time `json:"lock_time"   orm:"lock_time"   ` //
	CreateBy    string      `json:"create_by"   orm:"create_by"   ` //
	CreateAt    *gtime.Time `json:"create_at"   orm:"create_at"   ` //
	UpdateAt    *gtime.Time `json:"update_at"   orm:"update_at"   ` //
	DeleteAt    *gtime.Time `json:"delete_at"   orm:"delete_at"   ` //
}
