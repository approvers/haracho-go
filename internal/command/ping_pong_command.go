package command

import (
	"haracho-go/internal/client"
	"haracho-go/internal/service"
)

func init() {
	help := client.HelpContext{
		Name:            "!ping",
		Description:     "PingPong",
		ArgsDescription: nil,
	}

	service.GetCommandCollection().AddCommand(&help, "!ping", func(arg *client.Arg, ctx client.CommandContext) {
		ctx.SendMessage("pong!")
	})
}
