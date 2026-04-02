// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// Resource is the golang structure for table resource.
type Resource struct {
	ResourceUuid uuid.UUID   `json:"resource_uuid" orm:"resource_uuid" ` //
	CreateAt     *gtime.Time `json:"create_at"     orm:"create_at"     ` //
	UpdateAt     *gtime.Time `json:"update_at"     orm:"update_at"     ` //
	DeleteAt     *gtime.Time `json:"delete_at"     orm:"delete_at"     ` //
	Path         string      `json:"path"          orm:"path"          ` //
	Type         string      `json:"type"          orm:"type"          ` //
}
