// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Judger is the golang structure of table judger for DAO operations like Where/Data.
type Judger struct {
	g.Meta   `orm:"table:judger, do:true"`
	Jid      interface{} //
	Status   interface{} //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
