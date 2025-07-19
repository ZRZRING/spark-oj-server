// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Content is the golang structure for table content.
type Content struct {
	Cid      string      `json:"cid"       orm:"cid"       ` //
	Type     int         `json:"type"      orm:"type"      ` //
	Payload  string      `json:"payload"   orm:"payload"   ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at" ` //
}
