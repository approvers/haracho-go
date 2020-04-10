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

func NewClient(builder func(collection *client.CommandCollection)) *Client {
	collection := new(client.CommandCollection)
	builder(collection)
	return &Client{Collection: collection}
}

func (c *Client) Parse() (*arg.Arg, client.CommandContext) {
	ctx := new(CommandContext)
	c.context = ctx
	return arg.NewArg(c.input), ctx
}

func (c *Client) Execute(input string) string {
	c.input = input
	client.ExecuteCommand(c, c.Collection)
	return c.context.result
}
