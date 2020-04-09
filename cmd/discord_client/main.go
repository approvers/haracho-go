package main

import (
	"haracho-go/internal/client/discord"
	"os"
)

func main() {
	c := discord.Client{Token: os.Getenv("TOKEN")}
	c.Start()
}
