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
	Sid        any         //
	Title      any         //
	Pid        any         //
	Username   any         //
	Cid        any         //
	Result     any         //
	Language   any         //
	MemoryCost any         //
	TimeCost   any         //
	Code       any         //
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
	DeleteAt   *gtime.Time //
}
