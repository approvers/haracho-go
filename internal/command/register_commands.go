package command

import "haracho-go/internal/command/pingpong"

func RegisterCommands() {
	pingpong.Init()
}
