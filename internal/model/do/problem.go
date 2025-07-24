// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Problem is the golang structure of table problem for DAO operations like Where/Data.
type Problem struct {
	g.Meta      `orm:"table:problem, do:true"`
	Pid         interface{} //
	Title       interface{} //
	JudgeType   interface{} //
	TimeLimit   interface{} //
	MemoryLimit interface{} //
	CreateBy    interface{} //
	Rating      interface{} //
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
}
