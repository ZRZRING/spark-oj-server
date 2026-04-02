// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Resource is the golang structure of table resource for DAO operations like Where/Data.
type Resource struct {
	g.Meta       `orm:"table:resource, do:true"`
	ResourceUuid any         //
	CreateAt     *gtime.Time //
	UpdateAt     *gtime.Time //
	DeleteAt     *gtime.Time //
	Path         any         //
	Type         any         //
}
