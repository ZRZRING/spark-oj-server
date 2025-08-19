package service

type RunReq struct {
	Content     string
	TimeLimit   int
	MemeryLimit int
	StdIn       string
}

type RunRes struct {
	Status string
	StdOut string
}

func RunCode(req *RunReq) (res *RunRes, err error) {
	res = &RunRes{}

	return res, nil
}
