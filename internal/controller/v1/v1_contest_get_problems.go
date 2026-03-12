package v1

import (
	"context"
	"encoding/json"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetProblems(ctx context.Context, req *contest.GetProblemsReq) (res *contest.GetProblemsRes, err error) {
	res = &contest.GetProblemsRes{}

	result, err := dao.Contest.Ctx(ctx).
		Value("problems", "cid", req.Cid)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	problemIds := make([]int, 0)
	err = json.Unmarshal([]byte(result.String()), &problemIds)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	problems := make([]contest.GetProblemsItem, 0)
	for _, problemId := range problemIds {
		problem := &entity.Problem{}
		err = dao.Problem.Ctx(ctx).
			Where("pid", problemId).
			Scan(problem)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}

		problems = append(problems, contest.GetProblemsItem{
			Pid:       gconv.String(problem.Pid),
			Title:     problem.Title,
			JudgeType: problem.JudgeType,
		})
	}

	res = &contest.GetProblemsRes{
		Total:    len(problems),
		Problems: problems,
	}

	return res, nil
}
