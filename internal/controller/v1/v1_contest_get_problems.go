package v1

import (
	"context"
	"encoding/json"
	"spark-oj/internal/dao"
	"spark-oj/internal/model/entity"
	"spark-oj/pkg/enums"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj/api/v1/contest"
)

type contestProblemCountRow struct {
	ProblemId   int `json:"problem_id"`
	SubmitCount int `json:"submit_count"`
	AcceptCount int `json:"accept_count"`
}

func (c *ControllerContest) GetProblems(ctx context.Context, req *contest.GetProblemsReq) (res *contest.GetProblemsRes, err error) {
	res = &contest.GetProblemsRes{}

	result, err := dao.Contest.Ctx(ctx).
		Value("problems", "contest_id", req.ContestId)
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
			Where("problem_id", problemId).
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

	if len(problemIds) > 0 {
		var counts []contestProblemCountRow
		err = g.DB().Ctx(ctx).Raw(`
			SELECT problem_id,
				COUNT(*) AS submit_count,
				COUNT(*) FILTER (WHERE result = 'Accepted') AS accept_count
			FROM submission
			WHERE contest_id = ? AND problem_id IN (?) AND delete_at IS NULL
			GROUP BY problem_id
		`, req.ContestId, problemIds).Scan(&counts)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		countMap := make(map[int]*contestProblemCountRow, len(counts))
		for i := range counts {
			countMap[counts[i].ProblemId] = &counts[i]
		}
		for i, p := range problems {
			if c, ok := countMap[gconv.Int(p.ProblemId)]; ok {
				problems[i].SubmitCount = c.SubmitCount
				problems[i].AcceptCount = c.AcceptCount
			}
		}
	}

	res = &contest.GetProblemsRes{
		Total:    len(problems),
		Problems: problems,
	}

	return res, nil
}
