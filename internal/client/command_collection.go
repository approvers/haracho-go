package client

import "haracho-go/internal/client/arg"

type CommandCollection struct {
	commands []CommandExecutable
}

func (c CommandCollection) execute(arg *arg.Arg, ctx CommandContext) {
	for _, v := range c.commands {
		v.execute(arg, ctx)
	}
}

func (c *CommandCollection) AddCommand(command CommandExecutable) {
	c.commands = append(c.commands, command)
}

func (c *CommandCollection) AddPrefixCommand(help *HelpContext, prefix string, processor func(arg *arg.Arg, ctx CommandContext)) {
	command := &prefixCommand{
		Help:      help,
		prefix:    prefix,
		processor: processor,
	}

	c.commands = append(c.commands, command)
}
