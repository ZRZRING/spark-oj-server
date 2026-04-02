// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Problem is the golang structure for table problem.
type Problem struct {
	ProblemId   int         `json:"problem_id"   orm:"problem_id"   ` //
	CreateAt    *gtime.Time `json:"create_at"    orm:"create_at"    ` //
	UpdateAt    *gtime.Time `json:"update_at"    orm:"update_at"    ` //
	DeleteAt    *gtime.Time `json:"delete_at"    orm:"delete_at"    ` //
	Title       string      `json:"title"        orm:"title"        ` //
	JudgeType   string      `json:"judge_type"   orm:"judge_type"   ` //
	TimeLimit   int         `json:"time_limit"   orm:"time_limit"   ` //
	MemoryLimit int         `json:"memory_limit" orm:"memory_limit" ` //
	CreateBy    string      `json:"create_by"    orm:"create_by"    ` //
	Rating      int         `json:"rating"       orm:"rating"       ` //
	Content     string      `json:"content"      orm:"content"      ` //
}
