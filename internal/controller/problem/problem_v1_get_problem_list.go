package problem

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/problem/v1"
)

func (c *ControllerV1) GetProblemList(ctx context.Context, req *v1.GetProblemListReq) (res *v1.GetProblemListRes, err error) {
	res = &v1.GetProblemListRes{}
	md := dao.Problem.Ctx(ctx)

	// 获取分页信息
	problems := make([]*entity.Problem, 0)
	tot := new(int)
	err = md.OrderAsc("pid").Page(req.Page, req.Size).ScanAndCount(problems, tot, false)
	if err == nil {
		return nil, err
	}

	// 处理返回信息
	err = gconv.Scan(problems, res)
	if err != nil {
		return nil, err
	}
	res.Total = *tot

	return res, nil
}
