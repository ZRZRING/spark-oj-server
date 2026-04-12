package v1

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"spark-oj/api/v1/healthy"
)

func (c *ControllerHealthy) Check(ctx context.Context, req *healthy.CheckReq) (res *healthy.CheckRes, err error) {
	res = &healthy.CheckRes{}

	res.Database = checkDatabase(ctx)
	res.Redis = checkRedis(ctx)
	res.GoJudge = checkGoJudge(ctx)

	dbLink, _ := g.Cfg().Get(ctx, "database.default.link")
	judgeUrl, _ := g.Cfg().Get(ctx, "judge.apiBaseUrl")
	redisAddr, _ := g.Cfg().Get(ctx, "redis.default.address")
	res.Config = &healthy.ConfigInfo{
		DatabaseLink: maskPassword(dbLink.String()),
		JudgeApiUrl:  judgeUrl.String(),
		RedisAddress: redisAddr.String(),
	}

	allOk := res.Database.Status == "ok" && res.Redis.Status == "ok" && res.GoJudge.Status == "ok"
	anyDown := res.Database.Status == "error" || res.Redis.Status == "error" || res.GoJudge.Status == "error"

	switch {
	case anyDown:
		res.Status = "down"
	case allOk:
		res.Status = "ok"
	default:
		res.Status = "degraded"
	}

	return
}

func checkDatabase(ctx context.Context) *healthy.ComponentStatus {
	start := time.Now()
	_, err := g.DB().Ctx(ctx).Raw("SELECT 1").Value()
	latency := time.Since(start)
	if err != nil {
		return &healthy.ComponentStatus{
			Status:  "error",
			Latency: latency.String(),
			Error:   err.Error(),
		}
	}
	return &healthy.ComponentStatus{
		Status:  "ok",
		Latency: latency.String(),
	}
}

func checkRedis(ctx context.Context) *healthy.ComponentStatus {
	start := time.Now()
	_, err := g.Redis().Do(ctx, "PING")
	latency := time.Since(start)
	if err != nil {
		return &healthy.ComponentStatus{
			Status:  "error",
			Latency: latency.String(),
			Error:   err.Error(),
		}
	}
	return &healthy.ComponentStatus{
		Status:  "ok",
		Latency: latency.String(),
	}
}

func checkGoJudge(ctx context.Context) *healthy.ComponentStatus {
	judgeUrl, _ := g.Cfg().Get(ctx, "judge.apiBaseUrl")
	if judgeUrl.IsEmpty() {
		return &healthy.ComponentStatus{
			Status: "error",
			Error:  "judge.apiBaseUrl not configured",
		}
	}

	start := time.Now()
	resp, err := g.Client().Get(ctx, judgeUrl.String())
	latency := time.Since(start)
	if err != nil {
		return &healthy.ComponentStatus{
			Status:  "error",
			Latency: latency.String(),
			Error:   err.Error(),
		}
	}
	_ = resp.Close()
	return &healthy.ComponentStatus{
		Status:  "ok",
		Latency: latency.String(),
	}
}

func maskPassword(dsn string) string {
	atIdx := -1
	colonIdx := -1
	tcpIdx := -1
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '@' && atIdx < 0 {
			atIdx = i
		}
		if dsn[i] == ':' && atIdx < 0 && colonIdx < 0 {
			colonIdx = i
		}
		if dsn[i:i+4] == "tcp(" && tcpIdx < 0 {
			tcpIdx = i
		}
	}
	if atIdx > 0 && colonIdx > 0 && tcpIdx > 0 {
		return dsn[:colonIdx] + ":****" + dsn[atIdx:]
	}
	return dsn
}
