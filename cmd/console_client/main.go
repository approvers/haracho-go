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
	client.ExecuteCommand(c, collection)
}
