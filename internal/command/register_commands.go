package command

import (
	"haracho-go/internal/client"
	"haracho-go/internal/command/pingpong"
)

func RegisterCommands(collection *client.CommandCollection) {
	pingpong.Build(collection)
}
