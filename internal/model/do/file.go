// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure of table file for DAO operations like Where/Data.
type File struct {
	g.Meta   `orm:"table:file, do:true"`
	Uuid     interface{} //
	Path     interface{} //
	Type     interface{} //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
