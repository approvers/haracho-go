package pingpong

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/test"
	"testing"
)

func TestPingPong(t *testing.T) {
	collection := new(client.CommandCollection)
	c := test.Client{Collection: collection}
	Build(collection)

	expect := "pong!"
	res := c.Execute("!ping")
	if res != "pong!" {
		t.Error("\nコマンド: !ping", "\n理想: ", expect, "\n実際: ", res)
	}
	t.Log("!ping 正常終了")
}
