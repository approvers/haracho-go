package ping_pong

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
	"haracho-go/internal/service"
)

func Init() {
	help := client.HelpContext{
		Name:            "!ping",
		Description:     "PingPong",
		ArgsDescription: nil,
	}

	service.GetCommandCollection().AddCommand(&help, "!ping", func(arg *arg.Arg, ctx client.CommandContext) {
		ctx.SendMessage("pong!")
	})
}
