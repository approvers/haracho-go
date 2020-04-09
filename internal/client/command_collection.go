package client

type CommandCollection struct {
	commands []CommandExecutable
}

func (c CommandCollection) execute(arg *Arg, ctx CommandContext) {
	for _, v := range c.commands {
		v.execute(arg, ctx)
	}
}

func (c *CommandCollection) AddCommand(help *HelpContext, arg string, processor func(arg *Arg, ctx CommandContext)) {
	command := command{
		Help:        help,
		commandName: arg,
		processor:   processor,
	}
	c.commands = append(c.commands, command)
}

func (c *CommandCollection) AddSubCommands(help *HelpContext, arg string, subCommands []CommandExecutable) {
	command := subCommand{
		Help:     help,
		arg:      arg,
		commands: subCommands,
	}
	c.commands = append(c.commands, command)
}
