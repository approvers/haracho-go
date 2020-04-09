package client

type subCommand struct {
	Help     *HelpContext
	arg      string
	commands []CommandExecutable
}

func (c *subCommand) ShouldExecute(arg *Arg) bool {
	return arg.Next().IsSameCommand(c.arg)
}

func (c subCommand) execute(arg *Arg, ctx CommandContext) {
	if !c.ShouldExecute(arg) {
		return
	}

	for _, v := range c.commands {
		v.execute(arg.Next(), ctx)
	}
}
