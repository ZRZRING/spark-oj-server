package content

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta  `path:"/content/{uuid}" method:"PUT" tags:"content" summary:"更新题面"`
	Uuid    string `p:"uuid" v:"required#题面标识不能为空" dc:"题面唯一标识"`
	Type    *int   `p:"type" v:"in:0,1#类型只能是0(md)或1(html)" dc:"0:md, 1:html，不传则不修改"`
	Content string `p:"content" dc:"题面内容，不传则不修改"`
}

type UpdateRes struct{}