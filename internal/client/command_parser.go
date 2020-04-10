package client

import "haracho-go/internal/client/arg"

type CommandParser interface {
	Parse() (*arg.Arg, CommandContext)
}
