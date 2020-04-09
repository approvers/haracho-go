package ping_pong

import (
	"haracho-go/internal/client"
	"haracho-go/internal/service"
)

func Init() {
	help := client.HelpContext{
		Name:            "!ping",
		Description:     "PingPong",
		ArgsDescription: nil,
	}

	service.GetCommandCollection().AddCommand(&help, "!ping", func(arg *client.Arg, ctx client.CommandContext) {
		ctx.SendMessage("pong!")
	})
}
