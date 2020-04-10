package console

import "fmt"

type CommandContext struct {
}

func (c CommandContext) SendMessage(message string) {
	fmt.Println(message)
}
