package client

type CommandParser interface {
	Parse() (*Arg, CommandContext)
}
