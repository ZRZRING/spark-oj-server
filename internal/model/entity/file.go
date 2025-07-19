// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	Path     string      `json:"path"      orm:"path"      ` //
	Type     string      `json:"type"      orm:"type"      ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at" ` //
}
