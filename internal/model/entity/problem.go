// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Problem is the golang structure for table problem.
type Problem struct {
	Pid         string      `json:"pid"          orm:"pid"          ` //
	Title       string      `json:"title"        orm:"title"        ` //
	JudgeType   int         `json:"judge_type"   orm:"judge_type"   ` //
	Time        int         `json:"time"         orm:"time"         ` //
	Memory      int         `json:"memory"       orm:"memory"       ` //
	ContentType int         `json:"content_type" orm:"content_type" ` //
	Content     string      `json:"content"      orm:"content"      ` //
	CreateBy    string      `json:"create_by"    orm:"create_by"    ` //
	Rating      int         `json:"rating"       orm:"rating"       ` //
	Config      string      `json:"config"       orm:"config"       ` //
	CreateAt    *gtime.Time `json:"create_at"    orm:"create_at"    ` //
	UpdateAt    *gtime.Time `json:"update_at"    orm:"update_at"    ` //
	DeleteAt    *gtime.Time `json:"delete_at"    orm:"delete_at"    ` //
}
