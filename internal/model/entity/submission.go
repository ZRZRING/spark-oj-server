// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Submission is the golang structure for table submission.
type Submission struct {
	SubmissionId int         `json:"submission_id" orm:"submission_id" ` //
	CreateAt     *gtime.Time `json:"create_at"     orm:"create_at"     ` //
	UpdateAt     *gtime.Time `json:"update_at"     orm:"update_at"     ` //
	DeleteAt     *gtime.Time `json:"delete_at"     orm:"delete_at"     ` //
	ProblemId    int         `json:"problem_id"    orm:"problem_id"    ` //
	ContestId    int         `json:"contest_id"    orm:"contest_id"    ` //
	Username     string      `json:"username"      orm:"username"      ` //
	Result       string      `json:"result"        orm:"result"        ` //
	Language     string      `json:"language"      orm:"language"      ` //
	MemoryCost   int         `json:"memory_cost"   orm:"memory_cost"   ` //
	TimeCost     int         `json:"time_cost"     orm:"time_cost"     ` //
	Code         string      `json:"code"          orm:"code"          ` //
}
