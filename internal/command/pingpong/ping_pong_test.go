package pingpong

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/test"
	"haracho-go/internal/command"
	"testing"
)

func TestPingPong(t *testing.T) {
	collection := new(client.CommandCollection)
	c := test.Client{Collection: collection}
	command.RegisterCommands(collection)

	expect := "pong!"
	res := c.Execute("!ping")
	if res != "pong!" {
		t.Error("\nコマンド: !ping", "\n理想: ", expect, "\n実際: ", res)
	}
	t.Log("!ping 正常終了")
}
