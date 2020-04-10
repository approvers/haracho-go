package client

import "haracho-go/internal/client/arg"

type subCommand struct {
	Help     *HelpContext
	arg      string
	commands []CommandExecutable
}

func (c *subCommand) ShouldExecute(arg *arg.Arg) bool {
	return arg.Next().IsSameCommand(c.arg)
}

func (c *subCommand) execute(arg *arg.Arg, ctx CommandContext) {
	if !c.ShouldExecute(arg) {
		return
	}

	for _, v := range c.commands {
		v.execute(arg.Next(), ctx)
	}
}
