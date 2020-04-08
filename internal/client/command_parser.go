package client

type CommandParser interface {
	parse() Command
}
