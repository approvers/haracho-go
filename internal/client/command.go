package client

type Command interface {
	execute() *ExecResult
}

type ExecResult struct {
	Message string
}
