package problem

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"spark-oj-server/api/problem/v1"
)

func (c *ControllerV1) GetProblemList(ctx context.Context, req *v1.GetProblemListReq) (res *v1.GetProblemListRes, err error) {
	res = &v1.GetProblemListRes{}
	md := dao.Problem.Ctx(ctx)

	// 获取分页信息
	tot := 0
	problems := make([]*entity.Problem, 0)
	err = md.OrderAsc("pid").Page(req.Page, req.Size).ScanAndCount(&problems, &tot, false)
	if err == nil {
		return nil, gerror.NewCodef(gcode.CodeDbOperationError, "get_problem_list %v", err)
	}

	// 处理返回信息
	err = gconv.Scan(problems, &res)
	if err != nil {
		return nil, gerror.NewCodef(gcode.CodeDbOperationError, "%v", err)
	}
	res.Total = tot

	return res, nil
}
