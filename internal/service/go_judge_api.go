package service

import (
	"spark-oj-server/pkg/enums"
)

// 默认资源限制
var (
	DefaultCPULimit    = int64(10000000000) // 10秒
	DefaultMemoryLimit = int64(256000000)   // 256MB
	DefaultProcLimit   = 10
	DefaultOutputLimit = 10240 // 10KB
)

// 单份代码运行请求
type ExecuteCodeRequest struct {
	Code     string         `json:"code"`     // 源代码
	Input    string         `json:"input"`    // 输入数据
	Language enums.Language `json:"language"` // 编程语言 (cpp, c, java, python等)
}

// 单份代码运行响应
type ExecuteCodeResponse struct {
	Status     enums.JudgeStatus `json:"status"`     // 判题状态
	Output     string            `json:"output"`     // 程序输出
	Error      string            `json:"error"`      // 错误信息
	Time       int64             `json:"time"`       // 运行时间(纳秒)
	Memory     int64             `json:"memory"`     // 内存使用(字节)
	ExitStatus int64             `json:"exitStatus"` // 退出状态码
}

// go-judge API 请求结构
type GoJudgeRequest struct {
	Cmd []GoJudgeCmd `json:"cmd"`
}

// go-judge API 响应结构
type GoJudgeResponse []GoJudgeResult

// go-judge 命令结构
type GoJudgeCmd struct {
	Args          []string               `json:"args"`                    // 命令参数
	Env           []string               `json:"env"`                     // 环境变量
	Files         []GoJudgeFile          `json:"files"`                   // 传入文件
	CPULimit      int64                  `json:"cpuLimit"`                // CPU 时间限制(纳秒)
	MemoryLimit   int64                  `json:"memoryLimit"`             // 内存限制(字节)
	ProcLimit     int                    `json:"procLimit"`               // 进程限制
	CopyIn        map[string]interface{} `json:"copyIn"`                  // 传入文件内容
	CopyOut       []string               `json:"copyOut,omitempty"`       // 输出文件列表
	CopyOutCached []string               `json:"copyOutCached,omitempty"` // 缓存输出文件列表
}

// go-judge 文件接口
type GoJudgeFile any

// GoJudgeFile 支持类型

// LocalFile 本地文件
type LocalFile struct {
	Src string `json:"src"` // 文件绝对路径
}

// MemoryFile 内存文件
type MemoryFile struct {
	Content string `json:"content"` // 文件内容
}

// Collector 收集器
type Collector struct {
	Name string `json:"name"` // copyOut 文件名
	Max  int    `json:"max"`  // 最大大小限制
}

// go-judge 响应结构

type GoJudgeResult struct {
	Status     enums.JudgeStatus `json:"status"`
	Error      string            `json:"error,omitempty"`
	ExitStatus int64             `json:"exitStatus"`
	Time       int64             `json:"time"`
	Memory     int64             `json:"memory"`
	RunTime    int64             `json:"runTime"`
	Files      map[string]string `json:"files,omitempty"`
	FileIds    map[string]string `json:"fileIds,omitempty"`
}

// 编程语言配置
type LanguageConfig struct {
	CompileCmd     []string // 编译命令
	RunCmd         []string // 运行命令
	SourceFileName string   // 源文件名
	ExecutableName string   // 可执行文件名
}

// 支持的编程语言配置
var LanguageConfigs = map[enums.Language]LanguageConfig{
	"cpp": {
		CompileCmd:     []string{"/usr/bin/g++", "main.cpp", "-o", "main"},
		RunCmd:         []string{"main"},
		SourceFileName: "main.cpp",
		ExecutableName: "main",
	},
	"c": {
		CompileCmd:     []string{"/usr/bin/gcc", "main.c", "-o", "main"},
		RunCmd:         []string{"main"},
		SourceFileName: "main.c",
		ExecutableName: "main",
	},
	"java": {
		CompileCmd:     []string{"/usr/bin/javac", "Main.java"},
		RunCmd:         []string{"/usr/bin/java", "Main"},
		SourceFileName: "Main.java",
		ExecutableName: "Main.class",
	},
	"python": {
		CompileCmd:     nil,
		RunCmd:         []string{"/usr/bin/python3", "main.py"},
		SourceFileName: "main.py",
		ExecutableName: "",
	},
}
