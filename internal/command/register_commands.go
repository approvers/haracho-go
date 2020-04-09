package command

import "haracho-go/internal/command/ping_pong"

func RegisterCommands() {
	ping_pong.Init()
}
