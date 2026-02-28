package content

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta `path:"/content" method:"POST" tags:"content" summary:"创建题面"`
	Pid    int    `p:"pid" v:"required#题目ID不能为空" dc:"题目ID"`
	Type   int    `p:"type" v:"required|in:0,1#类型不能为空|类型只能是0(md)或1(html)" dc:"0:md, 1:html"`
	Content string `p:"content" v:"required#题面内容不能为空" dc:"题面内容，md或html格式"`
}

type CreateRes struct {
	Uuid string `json:"uuid" dc:"题面唯一标识"`
}