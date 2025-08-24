package service

type RunReq struct {
	Code        string
	TimeLimit   int
	MemeryLimit int
	StdIn       string
}

type RunRes struct {
	Status string
	StdOut string
}

func CompileCode(code string) (err error) {
	return nil
}

func RunCode(req *RunReq) (res *RunRes, err error) {
	res = &RunRes{}

	return res, nil
}
