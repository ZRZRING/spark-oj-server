package service

type JudgeReq struct {
	Content     string
	TimeLimit   int
	MemeryLimit int
	StdIn       string
	StdOut      string
}

type JudgeRes struct {
	Status string
}

type RunReq struct {
	Content     string
	TimeLimit   int
	MemeryLimit int
	StdIn       string
}

type RunRes struct {
	StdOut string
}

func judge() {

}
