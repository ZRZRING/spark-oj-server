package enums

// 判题结果状态枚举
type JudgeResultStatus string

const (
	JudgeStatusAccepted            JudgeResultStatus = "Accepted"
	JudgeStatusWrongAnswer         JudgeResultStatus = "Wrong Answer"
	JudgeStatusTimeLimitExceeded   JudgeResultStatus = "Time Limit Exceeded"
	JudgeStatusMemoryLimitExceeded JudgeResultStatus = "Memory Limit Exceeded"
	JudgeStatusRuntimeError        JudgeResultStatus = "Runtime Error"
	JudgeStatusCompileError        JudgeResultStatus = "Compile Error"
)

// 编程语言枚举
type Language string

const (
	LanguageC      Language = "c"
	LanguageCPP    Language = "cpp"
	LanguageJava   Language = "java"
	LanguagePython Language = "python"
)
