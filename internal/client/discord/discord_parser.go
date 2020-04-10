package discord

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
)

type Parser struct {
	message string
	ctx     client.CommandContext
}

func (p Parser) Parse() (*arg.Arg, client.CommandContext) {
	return arg.NewArg(p.message), p.ctx
}
