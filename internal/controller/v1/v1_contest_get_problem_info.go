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
	err = dao.Contest.Ctx(ctx).Where("cid", req.Cid).Scan(contestEntity)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if contestEntity.Cid == 0 {
		return nil, gerror.New("比赛不存在")
	}

	// 2. 解析比赛题目列表，校验题目是否属于该比赛
	var problemIds []int
	err = json.Unmarshal([]byte(contestEntity.Problems), &problemIds)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	pidInt := gconv.Int(req.Pid)
	found := false
	for _, id := range problemIds {
		if id == pidInt {
			found = true
			break
		}
	}
	if !found {
		return nil, gerror.New("该题目不属于此比赛")
	}

	// 3. 获取题目信息
	problem := &entity.Problem{}
	err = dao.Problem.Ctx(ctx).Where("pid", req.Pid).Scan(problem)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if problem.Pid == 0 {
		return nil, gerror.New("题目不存在")
	}

	// 4. 构建响应
	res = &contest.GetProblemInfoRes{
		Pid:          gconv.String(problem.Pid),
		Title:        problem.Title,
		JudgeType:    enums.JudgeType(problem.JudgeType),
		TimeLimit:    problem.TimeLimit,
		MemoryLimit:  problem.MemoryLimit,
		Rating:       problem.Rating,
		Content:      problem.Content,
		Cid:          req.Cid,
		ContestTitle: contestEntity.Title,
	}

	return res, nil
}