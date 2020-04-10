package client

import "haracho-go/internal/client/arg"

type CommandExecutable interface {
	execute(arg *arg.Arg, ctx CommandContext)
}

type command struct {
	Help        *HelpContext
	commandName string
	processor   func(arg *arg.Arg, ctx CommandContext)
}

func NewCommand(help *HelpContext, commandName string, processor func(arg *arg.Arg, ctx CommandContext)) *command {
	return &command{
		Help:        help,
		commandName: commandName,
		processor:   processor,
	}
}

func (c *command) ShouldExecute(arg *arg.Arg) bool {
	return arg.IsSameCommand(c.commandName)
}

func (c *command) execute(arg *arg.Arg, ctx CommandContext) {
	if c.ShouldExecute(arg) {
		c.processor(arg, ctx)
	}
}
