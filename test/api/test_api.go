package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "http://127.0.0.1:8000"

var (
	passed       int
	failed       int
	token        string
	problemID    string
	contestID    string
	submissionID string
)

type APIResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func test(name string, fn func() error) {
	fmt.Printf("  %-50s ", name)
	if err := fn(); err != nil {
		failed++
		fmt.Printf("FAIL\n       -> %s\n", err)
	} else {
		passed++
		fmt.Println("PASS")
	}
}

func request(method, path string, body interface{}) (*APIResponse, error) {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	var apiResp APIResponse
	if err := json.Unmarshal(raw, &apiResp); err != nil {
		return nil, fmt.Errorf("status=%d, body=%s", resp.StatusCode, string(raw))
	}
	return &apiResp, nil
}

func expectCode(code int, resp *APIResponse) error {
	if resp.Code != code {
		return fmt.Errorf("expected code %d, got %d: %s", code, resp.Code, resp.Message)
	}
	return nil
}

func main() {
	start := time.Now()
	fmt.Println("=" + strings.Repeat("=", 59))
	fmt.Println(" SparkOJ API 接口测试")
	fmt.Println("=" + strings.Repeat("=", 59))
	fmt.Printf(" 目标: %s\n", baseURL)
	fmt.Println()

	// ===== 用户模块 =====
	fmt.Println("[用户模块]")
	test("POST /register - 注册测试用户", func() error {
		resp, err := request("POST", "/register", map[string]string{
			"username":    "apitest",
			"password":    "test123456",
			"re_password": "test123456",
		})
		if err != nil {
			return err
		}
		// 注册成功或用户已存在都算通过
		if resp.Code != 0 && resp.Code != 51 {
			return fmt.Errorf("code=%d msg=%s", resp.Code, resp.Message)
		}
		return nil
	})

	test("POST /register - 密码不一致", func() error {
		resp, err := request("POST", "/register", map[string]string{
			"username":    "apitest2",
			"password":    "test123456",
			"re_password": "different",
		})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("POST /register - 参数缺失", func() error {
		resp, err := request("POST", "/register", map[string]string{
			"username": "apitest3",
		})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("POST /login - 登录", func() error {
		resp, err := request("POST", "/login", map[string]string{
			"username": "apitest",
			"password": "test123456",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Token string `json:"token"`
		}
		if err := json.Unmarshal(resp.Data, &data); err != nil {
			return fmt.Errorf("parse token: %w", err)
		}
		if data.Token == "" {
			return fmt.Errorf("token is empty")
		}
		token = data.Token
		return nil
	})

	test("POST /login - 密码错误", func() error {
		resp, err := request("POST", "/login", map[string]string{
			"username": "apitest",
			"password": "wrongpassword",
		})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("GET /training/{username} - 获取训练数据", func() error {
		resp, err := request("GET", "/training/apitest", nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	// ===== 题目模块 =====
	fmt.Println()
	fmt.Println("[题目模块]")

	test("POST /problem - 创建题目", func() error {
		resp, err := request("POST", "/problem", map[string]interface{}{
			"title":        "A+B Problem",
			"judge_type":   "Standard",
			"time_limit":   1000,
			"memory_limit": 128,
			"rating":       800,
			"create_by":    "apitest",
			"content":      "# A+B Problem\n给定两个整数，求它们的和。",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			ProblemId string `json:"problem_id"`
		}
		if err := json.Unmarshal(resp.Data, &data); err != nil {
			return fmt.Errorf("parse problemId: %w", err)
		}
		problemID = data.ProblemId
		return nil
	})

	test("POST /problem - 参数缺失", func() error {
		resp, err := request("POST", "/problem", map[string]interface{}{
			"title": "No Type",
		})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("GET /problem/{id} - 获取题目详情", func() error {
		if problemID == "" {
			return fmt.Errorf("没有可用的 problemId")
		}
		resp, err := request("GET", "/problem/"+problemID, nil)
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Title string `json:"title"`
		}
		json.Unmarshal(resp.Data, &data)
		if data.Title != "A+B Problem" {
			return fmt.Errorf("title mismatch: %s", data.Title)
		}
		return nil
	})

	test("GET /problem/{id} - 不存在的题目", func() error {
		resp, err := request("GET", "/problem/999999", nil)
		if err != nil {
			return err
		}
		// 应该返回错误码（51 或其他非零）
		if resp.Code == 0 {
			return fmt.Errorf("expected error for non-existent problem")
		}
		return nil
	})

	test("GET /problems?page=1&size=10 - 分页获取题目", func() error {
		resp, err := request("GET", "/problems?page=1&size=10", nil)
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Total int `json:"total"`
		}
		json.Unmarshal(resp.Data, &data)
		if data.Total < 1 {
			return fmt.Errorf("expected total >= 1, got %d", data.Total)
		}
		return nil
	})

	test("GET /problems - 缺少分页参数", func() error {
		resp, err := request("GET", "/problems", nil)
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("PUT /problem/{id} - 更新题目", func() error {
		if problemID == "" {
			return fmt.Errorf("没有可用的 problemId")
		}
		resp, err := request("PUT", "/problem/"+problemID, map[string]interface{}{
			"title":        "A+B Problem V2",
			"judge_type":   "Standard",
			"time_limit":   2000,
			"memory_limit": 256,
		})
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	// ===== 代码运行模块 =====
	fmt.Println()
	fmt.Println("[代码运行模块]")

	test("POST /run - 运行C++代码", func() error {
		resp, err := request("POST", "/run", map[string]string{
			"code":     "#include <iostream>\nint main() { int a, b; std::cin >> a >> b; std::cout << a + b; return 0; }",
			"input":    "1 2",
			"language": "cpp",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Status string `json:"status"`
			Output string `json:"output"`
		}
		json.Unmarshal(resp.Data, &data)
		if data.Output != "3" {
			return fmt.Errorf("expected output '3', got '%s'", data.Output)
		}
		return nil
	})

	test("POST /run - 运行Python代码", func() error {
		resp, err := request("POST", "/run", map[string]string{
			"code":     "print(input())",
			"input":    "hello",
			"language": "python",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Output string `json:"output"`
		}
		json.Unmarshal(resp.Data, &data)
		if !strings.Contains(data.Output, "hello") {
			return fmt.Errorf("expected output containing 'hello', got '%s'", data.Output)
		}
		return nil
	})

	// ===== 评测模块 =====
	fmt.Println()
	fmt.Println("[评测模块]")

	test("POST /judge - 提交判题", func() error {
		if problemID == "" {
			return fmt.Errorf("没有可用的 problemId")
		}
		resp, err := request("POST", "/judge", map[string]string{
			"problem_id": problemID,
			"username":   "apitest",
			"code":       "#include <iostream>\nint main() { int a, b; std::cin >> a >> b; std::cout << a + b; return 0; }",
			"language":   "cpp",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			SubmissionId string `json:"submission_id"`
			Result       string `json:"result"`
		}
		json.Unmarshal(resp.Data, &data)
		submissionID = data.SubmissionId
		return nil
	})

	test("POST /judge - 缺少必填字段", func() error {
		resp, err := request("POST", "/judge", map[string]string{
			"username": "apitest",
		})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	// ===== 提交记录模块 =====
	fmt.Println()
	fmt.Println("[提交记录模块]")

	test("GET /submission/{id} - 获取提交详情", func() error {
		if submissionID == "" {
			return fmt.Errorf("没有可用的 submissionId（跳过）")
		}
		resp, err := request("GET", "/submission/"+submissionID, nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	test("GET /submissions?page=1&size=10 - 分页获取提交列表", func() error {
		resp, err := request("GET", "/submissions?page=1&size=10", nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	// ===== 比赛模块 =====
	fmt.Println()
	fmt.Println("[比赛模块]")

	test("POST /contest - 创建比赛", func() error {
		resp, err := request("POST", "/contest", map[string]interface{}{
			"title":       "API测试比赛",
			"password":    "test123",
			"problems":    []int{},
			"description": "测试比赛",
			"practice":    true,
			"create_by":   "apitest",
		})
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		// 尝试从响应获取 contestId，如果没有返回数据则从列表获取
		if resp.Data != nil {
			var data struct {
				ContestId string `json:"contest_id"`
			}
			json.Unmarshal(resp.Data, &data)
			contestID = data.ContestId
		}
		return nil
	})

	test("POST /contest - 缺少必填字段", func() error {
		resp, err := request("POST", "/contest", map[string]string{})
		if err != nil {
			return err
		}
		return expectCode(51, resp)
	})

	test("GET /contests?page=1&size=10 - 获取比赛列表", func() error {
		resp, err := request("GET", "/contests?page=1&size=10", nil)
		if err != nil {
			return err
		}
		if err := expectCode(0, resp); err != nil {
			return err
		}
		var data struct {
			Total    int `json:"total"`
			Contests []struct {
				ContestId string `json:"contest_id"`
			} `json:"contests"`
		}
		json.Unmarshal(resp.Data, &data)
		// 如果前面没拿到 contestID，从列表取
		if contestID == "" && len(data.Contests) > 0 {
			contestID = data.Contests[0].ContestId
		}
		return nil
	})

	test("GET /contest/{id} - 获取比赛详情", func() error {
		if contestID == "" {
			return fmt.Errorf("没有可用的 contestId（跳过）")
		}
		resp, err := request("GET", "/contest/"+contestID, nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	test("GET /contest/{id}/problems - 获取比赛题目", func() error {
		if contestID == "" {
			return fmt.Errorf("没有可用的 contestId（跳过）")
		}
		resp, err := request("GET", "/contest/"+contestID+"/problems", nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	test("GET /contest/{id}/ranking - 获取排行榜", func() error {
		if contestID == "" {
			return fmt.Errorf("没有可用的 contestId（跳过）")
		}
		resp, err := request("GET", "/contest/"+contestID+"/ranking", nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	test("GET /contest/{id}/submissions - 获取比赛提交", func() error {
		if contestID == "" {
			return fmt.Errorf("没有可用的 contestId（跳过）")
		}
		resp, err := request("GET", fmt.Sprintf("/contest/%s/submissions?page=1&size=10", contestID), nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	test("PUT /contest/{id} - 更新比赛", func() error {
		if contestID == "" {
			return fmt.Errorf("没有可用的 contestId（跳过）")
		}
		resp, err := request("PUT", "/contest/"+contestID, map[string]interface{}{
			"title":       "API测试比赛V2",
			"description": "更新后的测试比赛",
			"create_by":   "apitest",
		})
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	// ===== Admin 模块 =====
	fmt.Println()
	fmt.Println("[Admin 模块]")

	test("GET /admin/protected - 无Token访问", func() error {
		oldToken := token
		token = ""
		resp, err := request("GET", "/admin/protected", nil)
		token = oldToken
		if err != nil {
			return err
		}
		if resp.Code == 0 {
			return fmt.Errorf("应该拒绝无Token访问")
		}
		return nil
	})

	test("GET /admin/protected - 带Token访问", func() error {
		if token == "" {
			return fmt.Errorf("没有可用的 token（跳过）")
		}
		resp, err := request("GET", "/admin/protected", nil)
		if err != nil {
			return err
		}
		return expectCode(0, resp)
	})

	// ===== 结果汇总 =====
	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 59))
	total := passed + failed
	fmt.Printf(" 测试完成: %d/%d 通过, %d 失败  耗时: %v\n", passed, total, failed, time.Since(start).Round(time.Millisecond))
	fmt.Println("=" + strings.Repeat("=", 59))

	if failed > 0 {
		os.Exit(1)
	}
}
