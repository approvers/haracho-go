package main

import (
	"bufio"
	"haracho-go/internal/client"
	"haracho-go/internal/client/console"
	"haracho-go/internal/command"
	"os"
)

func main() {
	collection := new(client.CommandCollection)
	command.RegisterCommands(collection)

	c := console.Client{Scanner: bufio.NewScanner(os.Stdin)}
	for i := 0; i < 5; i++ {
		client.ExecuteCommand(c, collection)
	}
}
