package healthy

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CheckReq struct {
	g.Meta `path:"/health" method:"GET" tags:"healthy" summary:"健康检查"`
}

type CheckRes struct {
	Status   string           `json:"status" dc:"总体状态: ok / degraded / down"`
	Database *ComponentStatus `json:"database"`
	Redis    *ComponentStatus `json:"redis"`
	GoJudge  *ComponentStatus `json:"go_judge"`
	Config   *ConfigInfo      `json:"config"`
}

type ComponentStatus struct {
	Status  string `json:"status" dc:"ok / error"`
	Latency string `json:"latency,omitempty" dc:"响应延迟"`
	Error   string `json:"error,omitempty" dc:"错误信息"`
}

type ConfigInfo struct {
	DatabaseLink string `json:"database_link"`
	JudgeApiUrl  string `json:"judge_api_url"`
	RedisAddress string `json:"redis_address"`
}
