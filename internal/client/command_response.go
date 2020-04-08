package client

type CommandResponse interface {
	respond(result *ExecResult)
}
