package client

type CommandExecutable interface {
	execute(arg *Arg, ctx CommandContext)
}

type command struct {
	Help      *HelpContext
	arg       string
	processor func(arg *Arg, ctx CommandContext)
}

func (c *command) ShouldExecute(arg *Arg) bool {
	return arg.IsSame(c.arg)
}

func (c command) execute(arg *Arg, ctx CommandContext) {
	if c.ShouldExecute(arg) {
		c.processor(arg, ctx)
	}
}
