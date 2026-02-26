// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Announcement is the golang structure of table announcement for DAO operations like Where/Data.
type Announcement struct {
	g.Meta   `orm:"table:announcement, do:true"`
	Title    interface{} //
	Content  interface{} //
	CreateBy interface{} //
	CreateAt interface{} //
	UpdateAt interface{} //
	DeleteAt interface{} //
}
