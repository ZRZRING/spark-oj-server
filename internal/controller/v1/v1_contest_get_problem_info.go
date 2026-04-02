package v1

import (
	"context"
	"encoding/json"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
	"spark-oj-server/pkg/enums"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/contest"
)

func (c *ControllerContest) GetProblemInfo(ctx context.Context, req *contest.GetProblemInfoReq) (res *contest.GetProblemInfoRes, err error) {
	// 1. 获取比赛信息
	contestEntity := &entity.Contest{}
	err = dao.Contest.Ctx(ctx).Where("contestId", req.ContestId).Scan(contestEntity)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if contestEntity.ContestId == 0 {
		return nil, gerror.New("比赛不存在")
	}

	// 2. 解析比赛题目列表，校验题目是否属于该比赛
	var problemIds []int
	err = json.Unmarshal([]byte(contestEntity.Problems), &problemIds)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	problemIdInt := gconv.Int(req.ProblemId)
	found := false
	for _, id := range problemIds {
		if id == problemIdInt {
			found = true
			break
		}
	}
	if !found {
		return nil, gerror.New("该题目不属于此比赛")
	}

	// 3. 获取题目信息
	problem := &entity.Problem{}
	err = dao.Problem.Ctx(ctx).Where("problemId", req.ProblemId).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if problem.ProblemId == 0 {
		return nil, gerror.New("题目不存在")
	}

	// 4. 构建响应
	res = &contest.GetProblemInfoRes{
		ProblemId:    gconv.String(problem.ProblemId),
		Title:        problem.Title,
		JudgeType:    enums.JudgeType(problem.JudgeType),
		TimeLimit:    problem.TimeLimit,
		MemoryLimit:  problem.MemoryLimit,
		Rating:       problem.Rating,
		Content:      problem.Content,
		ContestId:    req.ContestId,
		ContestTitle: contestEntity.Title,
	}

	return res, nil
}
