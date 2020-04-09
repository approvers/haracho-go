package console

import "fmt"

type Context struct {
}

func (c Context) SendMessage(message string) {
	fmt.Println(message)
}
