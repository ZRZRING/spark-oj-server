// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Submission is the golang structure for table submission.
type Submission struct {
	Sid      int         `json:"sid"       orm:"sid"       ` //
	Pid      string      `json:"pid"       orm:"pid"       ` //
	Uid      string      `json:"uid"       orm:"uid"       ` //
	Cid      string      `json:"cid"       orm:"cid"       ` //
	Result   string      `json:"result"    orm:"result"    ` //
	Language string      `json:"language"  orm:"language"  ` //
	Memory   int         `json:"memory"    orm:"memory"    ` //
	Time     int         `json:"time"      orm:"time"      ` //
	Code     string      `json:"code"      orm:"code"      ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at" ` //
}
