package test

type CommandContext struct {
	result string
}

func (c *CommandContext) SendMessage(message string) {
	c.result = message
}
