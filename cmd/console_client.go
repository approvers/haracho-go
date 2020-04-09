package main

import (
	"bufio"
	"haracho-go/internal/client"
	"haracho-go/internal/client/console"
	"haracho-go/internal/service"
	"os"
)

func main() {
	c := console.Client{Scanner: bufio.NewScanner(os.Stdin)}
	for range make([]int, 5) {
		client.ExecuteCommand(c, service.GetCommandCollection())
		break
	}
}
