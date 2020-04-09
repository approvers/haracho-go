package main

import (
	"bufio"
	"haracho-go/internal/client"
	"haracho-go/internal/client/console"
	"haracho-go/internal/command"
	"haracho-go/internal/service"
	"os"
)

func main() {
	command.RegisterCommands()

	c := console.Client{Scanner: bufio.NewScanner(os.Stdin)}
	for i := 0; i < 5; i++ {
		client.ExecuteCommand(c, service.GetCommandCollection())
	}
}
