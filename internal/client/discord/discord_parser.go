package discord

import (
	"haracho-go/internal/client"
	"strings"
)

type Parser struct {
	rawMessage string
	ctx        client.CommandContext
}

func (p Parser) Parse() (*client.Arg, client.CommandContext) {
	split := strings.Split(p.rawMessage, " ")
	return &client.Arg{CommandName: split[0], Args: split[1:]}, p.ctx
}
