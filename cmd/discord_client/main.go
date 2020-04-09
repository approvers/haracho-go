package main

import (
	"haracho-go/internal/client/discord"
	"haracho-go/internal/command"
	"os"
)

func main() {
	command.RegisterCommands()

	c := discord.Client{Token: os.Getenv("TOKEN")}
	c.Start()
}
