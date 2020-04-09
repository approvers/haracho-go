package service

import "haracho-go/internal/client"

var commandCollectionInstance = new(client.CommandCollection)

func GetCommandCollection() *client.CommandCollection {
	return commandCollectionInstance
}
