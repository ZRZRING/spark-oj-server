package user

import "github.com/gogf/gf/v2/frame/g"

type GetPageReq struct {
	g.Meta `path:"/users" method:"GET" tags:"user" summary:"获取用户列表"`
	Page   int `p:"page"`
	Size   int `p:"size"`
}

type GetPageRes struct {
	Total int           `json:"total"`
	Users []GetPageItem `json:"users"`
}

type GetPageItem struct {
	Username   string `json:"username"`
	UserRole   string `json:"user_role"`
	CreateTime string `json:"create_time"`
	Rating     string `json:"rating"`
}
