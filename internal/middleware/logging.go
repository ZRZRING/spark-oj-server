package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Logging 访问日志中间件
func Logging(r *ghttp.Request) {
	// 记录开始时间
	startTime := time.Now()

	// 读取请求体（用于记录）
	var requestBody string
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			requestBody = string(bodyBytes)
			// 将读取后的内容重新设置到请求中，确保后续处理器能读取
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	// 调用下一个中间件/处理器
	r.Middleware.Next()

	// 构建日志信息
	logData := g.Map{
		"method":   r.Method,
		"path":     r.URL.Path,
		"query":    r.URL.RawQuery,
		"status":   r.Response.Status,
		"duration": fmt.Sprintf("%v", time.Since(startTime)),
		// "ip":         r.GetClientIp(),
		// "user_agent": r.Header.Get("User-Agent"),
		// "time":       startTime.Format("2006-01-02 15:04:05"),
	}

	// 请求体
	if requestBody != "" {
		if len(requestBody) > 100 {
			logData["request_body"] = requestBody[:100] + "...(truncated)"
		} else {
			logData["request_body"] = requestBody
		}
	}

	// 响应体 Buffer
	responseBody := r.Response.BufferString()
	if responseBody != "" {
		if len(responseBody) > 100 {
			logData["response_body"] = responseBody[:100] + "...(truncated)"
		} else {
			logData["response_body"] = responseBody
		}
	}

	// 记录日志
	g.Log().Infof(r.Context(), "Access Log: %+v", logData)
}
