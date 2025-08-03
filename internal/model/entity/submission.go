// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Submission is the golang structure for table submission.
type Submission struct {
	Sid        int         `json:"sid"         orm:"sid"         ` //
	Title      string      `json:"title"       orm:"title"       ` //
	Pid        string      `json:"pid"         orm:"pid"         ` //
	Username   string      `json:"username"    orm:"username"    ` //
	Cid        string      `json:"cid"         orm:"cid"         ` //
	Result     string      `json:"result"      orm:"result"      ` //
	Language   string      `json:"language"    orm:"language"    ` //
	MemoryCost int         `json:"memory_cost" orm:"memory_cost" ` //
	TimeCost   int         `json:"time_cost"   orm:"time_cost"   ` //
	Code       string      `json:"code"        orm:"code"        ` //
	CreateAt   *gtime.Time `json:"create_at"   orm:"create_at"   ` //
	UpdateAt   *gtime.Time `json:"update_at"   orm:"update_at"   ` //
	DeleteAt   *gtime.Time `json:"delete_at"   orm:"delete_at"   ` //
}
