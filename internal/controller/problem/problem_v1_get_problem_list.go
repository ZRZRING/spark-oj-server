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
	var problems []*entity.Problem
	err = md.
		OrderAsc("pid").
		Page(req.Page, req.Size).
		Scan(&problems)
	if err == nil {
		return nil, gerror.NewCodef(gcode.CodeDbOperationError, "get_problem_list %v", err)
	}
	err = gconv.Scan(problems, &res)
	tot, err := md.Count()
	if err != nil {
		return nil, gerror.NewCodef(gcode.CodeDbOperationError, "%v", err)
	}
	res.Total = tot
	return
}
