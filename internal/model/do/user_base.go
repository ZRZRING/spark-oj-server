// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure of table user_base for DAO operations like Where/Data.
type UserBase struct {
	g.Meta    `orm:"table:user_base, do:true"`
	Username  any         //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
	DeleteAt  *gtime.Time //
	Password  any         //
	Nickname  any         //
	Role      any         //
	Solved    any         //
	Submitted any         //
	Rating    any         //
	Extra     any         //
}
