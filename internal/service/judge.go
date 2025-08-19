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

func judge() {

}
