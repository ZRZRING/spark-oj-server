// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure for table user_base.
type UserBase struct {
	Username   string      `json:"username"   orm:"username"   ` //
	Password   string      `json:"password"   orm:"password"   ` //
	Nickname   string      `json:"nickname"   orm:"nickname"   ` //
	Role       string      `json:"role"       orm:"role"       ` //
	Solved     int         `json:"solved"     orm:"solved"     ` //
	Submitted  int         `json:"submitted"  orm:"submitted"  ` //
	Rating     int         `json:"rating"     orm:"rating"     ` //
	CfId       string      `json:"cf_id"      orm:"cf_id"      ` //
	AtcId      string      `json:"atc_id"     orm:"atc_id"     ` //
	Company    string      `json:"company"    orm:"company"    ` //
	Department string      `json:"department" orm:"department" ` //
	Major      string      `json:"major"      orm:"major"      ` //
	Class      string      `json:"class"      orm:"class"      ` //
	Email      string      `json:"email"      orm:"email"      ` //
	Tel        string      `json:"tel"        orm:"tel"        ` //
	Avatar     string      `json:"avatar"     orm:"avatar"     ` //
	CreateAt   *gtime.Time `json:"create_at"  orm:"create_at"  ` //
	UpdateAt   *gtime.Time `json:"update_at"  orm:"update_at"  ` //
	DeleteAt   *gtime.Time `json:"delete_at"  orm:"delete_at"  ` //
}
