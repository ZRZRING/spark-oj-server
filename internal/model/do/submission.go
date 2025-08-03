// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Submission is the golang structure of table submission for DAO operations like Where/Data.
type Submission struct {
	g.Meta     `orm:"table:submission, do:true"`
	Sid        interface{} //
	Title      interface{} //
	Pid        interface{} //
	Username   interface{} //
	Cid        interface{} //
	Result     interface{} //
	Language   interface{} //
	MemoryCost interface{} //
	TimeCost   interface{} //
	Code       interface{} //
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
	DeleteAt   *gtime.Time //
}
