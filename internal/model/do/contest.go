// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Contest is the golang structure of table contest for DAO operations like Where/Data.
type Contest struct {
	g.Meta      `orm:"table:contest, do:true"`
	ContestId   any         //
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
	Title       any         //
	Password    any         //
	Description any         //
	StartTime   *gtime.Time //
	EndTime     *gtime.Time //
	LockTime    *gtime.Time //
	CreateBy    any         //
	Practice    any         //
	Problems    any         //
}
