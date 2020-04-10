package test

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
)

type Client struct {
	input      string
	context    *CommandContext
	Collection *client.CommandCollection
}

func (c *Client) Parse() (*arg.Arg, client.CommandContext) {
	ctx := new(CommandContext)
	return arg.NewArg(c.input), ctx
}

func (c *Client) Execute(input string) string {
	c.input = input
	client.ExecuteCommand(c, c.Collection)
	return c.context.result
}
