package test

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
	"haracho-go/internal/command"
)

type Client struct {
	input      string
	context    *CommandContext
	collection *client.CommandCollection
}

func (c *Client) Init() {
	collection := new(client.CommandCollection)
	command.RegisterCommands(collection)
	c.collection = collection
}

func (c *Client) Parse() (*arg.Arg, client.CommandContext) {
	ctx := new(CommandContext)
	return arg.NewArg(c.input), ctx
}

func (c *Client) Execute(input string) string {
	c.input = input
	client.ExecuteCommand(c, c.collection)
	return c.context.result
}
