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
	Pid         any         //
	Title       any         //
	JudgeType   any         //
	TimeLimit   any         //
	MemoryLimit any         //
	CreateBy    any         //
	Rating      any         //
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
	Content     any         //
	Analysis    any         //
}
