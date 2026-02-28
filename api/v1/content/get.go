package content

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/content/{uuid}" method:"GET" tags:"content" summary:"获取单个题面"`
	Uuid   string `p:"uuid" v:"required#题面标识不能为空" dc:"题面唯一标识"`
}

type GetRes struct {
	Uuid     string `json:"uuid" dc:"题面唯一标识"`
	Type     int    `json:"type" dc:"类型: 0:md, 1:html"`
	Content  string `json:"content" dc:"题面内容"`
	CreateAt string `json:"create_at" dc:"创建时间"`
	UpdateAt string `json:"update_at" dc:"更新时间"`
}