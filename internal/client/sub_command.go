package client

import "haracho-go/internal/client/arg"

type subCommand struct {
	Help        *HelpContext
	commandName string
	commands    []CommandExecutable
}

func NewSubCommand(help *HelpContext, commandName string, commands []CommandExecutable) *subCommand {
	return &subCommand{
		Help:        help,
		commandName: commandName,
		commands:    commands,
	}
}

func (c *subCommand) ShouldExecute(arg *arg.Arg) bool {
	return arg.IsSameCommand(c.commandName)
}

func (c *subCommand) execute(arg *arg.Arg, ctx CommandContext) {
	if !c.ShouldExecute(arg) {
		return
	}

	for _, v := range c.commands {
		v.execute(arg.Next(), ctx)
	}
}
