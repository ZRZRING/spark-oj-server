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
	g.Meta     `orm:"table:user_base, do:true"`
	Username   interface{} //
	Password   interface{} //
	Nickname   interface{} //
	Role       interface{} //
	Solved     interface{} //
	Submitted  interface{} //
	Rating     interface{} //
	CfId       interface{} //
	AtcId      interface{} //
	Company    interface{} //
	Department interface{} //
	Major      interface{} //
	Class      interface{} //
	Email      interface{} //
	Tel        interface{} //
	Avatar     interface{} //
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
	DeleteAt   *gtime.Time //
}
