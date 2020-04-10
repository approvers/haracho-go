package main

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/discord"
	"haracho-go/internal/command"
	"os"
)

func main() {
	collection := new(client.CommandCollection)
	command.RegisterCommands(collection)

	c := discord.Client{Token: os.Getenv("TOKEN")}
	c.Start()
}
