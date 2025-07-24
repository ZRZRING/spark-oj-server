// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	Tag      string      `json:"tag"       orm:"tag"       ` //
	Approved bool        `json:"approved"  orm:"approved"  ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at" ` //
}
