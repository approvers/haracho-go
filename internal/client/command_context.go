package client

type CommandContext interface {
	SendMessage(message string)
}
