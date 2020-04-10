package client

import "haracho-go/internal/client/arg"

type prefixCommand struct {
	Help      *HelpContext
	prefix    string
	processor func(arg *arg.Arg, ctx CommandContext)
}

func NewPrefixCommand(help *HelpContext, prefix string, processor func(arg *arg.Arg, ctx CommandContext)) *prefixCommand {
	return &prefixCommand{
		Help:      help,
		prefix:    prefix,
		processor: processor,
	}
}

func (c *prefixCommand) ShouldExecute(arg *arg.Arg) bool {
	return arg.StartsWith(c.prefix)
}

func (c *prefixCommand) execute(arg *arg.Arg, ctx CommandContext) {
	if c.ShouldExecute(arg) {
		c.execute(arg, ctx)
	}
}
