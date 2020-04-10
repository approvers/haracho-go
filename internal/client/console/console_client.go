package console

import (
	"bufio"
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
	_ "haracho-go/internal/command"
)

type Client struct {
	Scanner *bufio.Scanner
}

func (c Client) Parse() (*arg.Arg, client.CommandContext) {
	c.Scanner.Scan()
	input := c.Scanner.Text()
	return arg.NewArg(input), CommandContext{}
}
