package service

import (
	"encoding/json"
	"fmt"
	"spark-oj/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

// ExecuteCode 执行单次代码（编译并运行）
func ExecuteCode(ctx g.Ctx, request *ExecuteCodeRequest) (*ExecuteCodeResponse, error) {
	// 验证编程语言是否支持
	langConfig, exists := LanguageConfigs[request.Language]
	if !exists {
		return &ExecuteCodeResponse{
			Status: "Unsupported Language",
			Error:  fmt.Sprintf("Unsupported language: %s", request.Language),
		}, nil
	}

	var executableFileId string

	// 如果需要编译，先进行编译
	if langConfig.CompileCmd != nil {
		compileResult, err := compileCode(ctx, request.Code, langConfig)
		if err != nil {
			return nil, err
		}

		// 检查编译是否成功
		if compileResult.Status != "Accepted" {
			return &ExecuteCodeResponse{
				Status:     "Compile Error",
				Error:      compileResult.Files["stderr"],
				ExitStatus: compileResult.ExitStatus,
			}, nil
		}

		// 获取编译后的可执行文件ID
		for fileName, fileId := range compileResult.FileIds {
			if fileName == langConfig.ExecutableName {
				executableFileId = fileId
				break
			}
		}
	}

	// 运行代码
	runResult, err := runCode(ctx, request.Code, executableFileId, request.Input, langConfig)
	if err != nil {
		return nil, err
	}

	// 构造响应
	response := &ExecuteCodeResponse{
		Status:     runResult.Status,
		Output:     runResult.Files["stdout"],
		Error:      runResult.Files["stderr"],
		Time:       runResult.Time,
		Memory:     runResult.Memory,
		ExitStatus: runResult.ExitStatus,
	}

	return response, nil
}

// compileCode 编译代码
func compileCode(ctx g.Ctx, code string, langConfig LanguageConfig) (*GoJudgeResult, error) {
	request := &GoJudgeRequest{
		Cmd: []GoJudgeCmd{
			{
				Args: langConfig.CompileCmd,
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []GoJudgeFile{
					&MemoryFile{Content: ""},
					&Collector{Name: "stdout", Max: DefaultOutputLimit},
					&Collector{Name: "stderr", Max: DefaultOutputLimit},
				},
				CPULimit:    DefaultCPULimit,
				MemoryLimit: DefaultMemoryLimit,
				ProcLimit:   DefaultProcLimit,
				CopyIn: map[string]interface{}{
					langConfig.SourceFileName: map[string]string{
						"content": code,
					},
				},
				CopyOut:       []string{"stdout", "stderr"},
				CopyOutCached: []string{langConfig.ExecutableName},
			},
		},
	}

	return sendGoJudgeRequest(ctx, request)
}

// runCode 运行代码
func runCode(ctx g.Ctx, code, executableFileId, input string, langConfig LanguageConfig) (*GoJudgeResult, error) {
	cmd := GoJudgeCmd{
		Args: langConfig.RunCmd,
		Env:  []string{"PATH=/usr/bin:/bin"},
		Files: []GoJudgeFile{
			&MemoryFile{Content: input},
			&Collector{Name: "stdout", Max: DefaultOutputLimit},
			&Collector{Name: "stderr", Max: DefaultOutputLimit},
		},
		CPULimit:    DefaultCPULimit,
		MemoryLimit: DefaultMemoryLimit,
		ProcLimit:   DefaultProcLimit,
		CopyIn:      make(map[string]interface{}),
	}

	// 如果有可执行文件ID，使用缓存的可执行文件（编译型语言）
	if executableFileId != "" {
		executableName := langConfig.ExecutableName
		cmd.CopyIn[executableName] = map[string]string{
			"fileId": executableFileId,
		}
	} else {
		// 对于解释型语言，直接复制源文件
		cmd.CopyIn[langConfig.SourceFileName] = map[string]string{
			"content": code,
		}
	}

	request := &GoJudgeRequest{
		Cmd: []GoJudgeCmd{cmd},
	}

	return sendGoJudgeRequest(ctx, request)
}

// sendGoJudgeRequest 发送运行代码请求到 go-judge
func sendGoJudgeRequest(ctx g.Ctx, request *GoJudgeRequest) (res *GoJudgeResult, err error) {
	// 将请求转换为JSON
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %v", err)
	}

	g.Log().Debug(ctx, "Judge request:", string(requestData))

	// 发送POST请求到go-judge
	judgeApiBaseUrl := utils.GetConfigString("judge.apiBaseUrl")
	r, err := g.Client().Post(ctx, judgeApiBaseUrl+"/run", requestData)
	if err != nil {
		return nil, fmt.Errorf("send request failed: %v", err)
	}
	defer func(r *gclient.Response) {
		err = r.Close()
		if err != nil {
			g.Log().Error(ctx, "close response failed:", err)
		}
	}(r)

	// 解析响应
	responseBody := r.ReadAllString()
	g.Log().Debug(ctx, "Judge response:", responseBody)

	var response GoJudgeResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("empty response")
	}

	return &response[0], nil
}
