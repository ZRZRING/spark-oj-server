package v1

import (
	"context"
	"encoding/json"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetProblems(ctx context.Context, req *contest.GetProblemsReq) (res *contest.GetProblemsRes, err error) {
	res = &contest.GetProblemsRes{}

	result, err := dao.Contest.Ctx(ctx).
		Value("problems", "contestId", req.ContestId)
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
			Where("problemId", problemId).
			Scan(problem)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}

		problems = append(problems, contest.GetProblemsItem{
			ProblemId: gconv.String(problem.ProblemId),
			Title:     problem.Title,
			JudgeType: enums.JudgeType(problem.JudgeType),
		})
	}

	res = &contest.GetProblemsRes{
		Total:    len(problems),
		Problems: problems,
	}

	return res, nil
}
