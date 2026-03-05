package consts

const (
	ProblemJudgeType_Unknown = 0
	ProblemJudgeType_ACM     = 1
)

var ProblemJudgeTypeMap = map[int]string{
	ProblemJudgeType_Unknown: "Unknown",
	ProblemJudgeType_ACM:     "ACM",
}

const (
	ContentType_Unknown  = 0
	ContentType_Markdown = 1
	ContentType_HTML     = 2
)

var ContentTypeMap = map[int]string{
	ContentType_Unknown:  "Unknown",
	ContentType_Markdown: "Markdown",
	ContentType_HTML:     "HTML",
}
