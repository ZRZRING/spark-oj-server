package codeforces

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserInfo struct {
	LastOnlineTimeSeconds gtime.Time `json:"lastOnlineTimeSeconds"`
	Rating                int        `json:"rating"`
	MaxRating             int        `json:"maxRating"`
}
