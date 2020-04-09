package console

import (
	"bufio"
	"haracho-go/internal/client"
	_ "haracho-go/internal/command"
	"strings"
)

type Client struct {
	Scanner *bufio.Scanner
}

func (c Client) Parse() (*client.Arg, client.CommandContext) {
	c.Scanner.Scan()
	input := c.Scanner.Text()
	split := strings.Split(input, " ")
	return &client.Arg{CommandName: split[0], Args: split[1:]}, CommandContext{}
}
