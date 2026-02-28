package content

import "github.com/gogf/gf/v2/frame/g"

type GetListReq struct {
	g.Meta `path:"/content" method:"GET" tags:"content" summary:"获取题面列表"`
	Pid    int    `p:"pid" v:"required#题目ID不能为空" dc:"题目ID"`
	Type   *int   `p:"type" dc:"类型筛选: 0:md, 1:html，不传则查询所有类型"`
	Page   int    `p:"page" d:"1" dc:"页码"`
	PageSize int  `p:"page_size" d:"10" dc:"每页数量"`
}

type ContentItem struct {
	Uuid     string `json:"uuid" dc:"题面唯一标识"`
	Type     int    `json:"type" dc:"类型: 0:md, 1:html"`
	Content  string `json:"content" dc:"题面内容"`
	CreateAt string `json:"create_at" dc:"创建时间"`
	UpdateAt string `json:"update_at" dc:"更新时间"`
}

type GetListRes struct {
	List  []ContentItem `json:"list" dc:"题面列表"`
	Total int           `json:"total" dc:"总数"`
	Page  int           `json:"page" dc:"当前页码"`
}