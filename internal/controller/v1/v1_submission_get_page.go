package v1

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"spark-oj-server/api/v1/submission"
	"spark-oj-server/internal/dao"
	"spark-oj-server/internal/model/entity"
)

func (c *ControllerSubmission) GetPage(ctx context.Context, req *submission.GetPageReq) (res *submission.GetPageRes, err error) {
	// 查询总数
	total, err := dao.Submission.Ctx(ctx).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询提交总数失败")
	}

	// 查询提交列表
	var submissions []*entity.Submission
	err = dao.Submission.Ctx(ctx).
		Page(req.Page, req.Size).
		Scan(&submissions)
	if err != nil {
		return nil, gerror.Wrap(err, "查询提交列表失败")
	}

	// 转换为返回格式
	res = &submission.GetPageRes{
		Total: total,
	}
	for _, sub := range submissions {
		res.Submissions = append(res.Submissions, &submission.GetPageItem{
			Sid:        gconv.String(sub.Sid),
			Pid:        sub.Pid,
			Cid:        sub.Cid,
			Username:   sub.Username,
			Result:     sub.Result,
			Language:   sub.Language,
			MemoryCost: fmt.Sprintf("%.2f", float64(sub.MemoryCost)/1024.0),
			TimeCost:   fmt.Sprintf("%d", sub.TimeCost),
			CreateTime: *sub.CreateAt,
		})
	}

	return res, nil
}
