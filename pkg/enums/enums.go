package enums

// 判题状态枚举
type JudgeStatus string

const (
	JudgeStatusWaiting             JudgeStatus = "Waiting"
	JudgeStatusRunning             JudgeStatus = "Running"
	JudgeStatusAccepted            JudgeStatus = "Accepted"
	JudgeStatusWrongAnswer         JudgeStatus = "Wrong Answer"
	JudgeStatusTimeLimitExceeded   JudgeStatus = "Time Limit Exceeded"
	JudgeStatusMemoryLimitExceeded JudgeStatus = "Memory Limit Exceeded"
	JudgeStatusRuntimeError        JudgeStatus = "Runtime Error"
	JudgeStatusCompileError        JudgeStatus = "Compile Error"
)

type RankingProblemStatus string

const (
	RankingStatusAccepted RankingProblemStatus = "Accepted"
	RankingStatusReject   RankingProblemStatus = "Reject"
	RankingStatusEmpty    RankingProblemStatus = ""
)

// 编程语言枚举
type Language string

const (
	LanguageC      Language = "c"
	LanguageCPP    Language = "cpp"
	LanguageJava   Language = "java"
	LanguagePython Language = "python"
)

type JudgeType string

const (
	JudgeTypeStandard      JudgeType = "Standard"
	JudgeTypeSpecialJudge  JudgeType = "Special Judge"
	JudgeTypeInteractive   JudgeType = "Interactive"
	JudgeTypeCommunication JudgeType = "Communication"
	JudgeTypeOutputOnly    JudgeType = "Output Only"
	JudgeTypeUnknown       JudgeType = "Unknown"
)
