package grpc_service

import (
	"fmt"
	"github.com/split-notes/pennant-admin-backend/configs"
	"google.golang.org/grpc"
)

type Connection struct {
	Server        *grpc.ClientConn
	GreeterClient GreeterClient
}

var server Connection

func Start(config configs.Configuration) (*Connection, error){
	address := fmt.Sprintf("%s:%s", config.FlaggerHost, config.FlaggerPort)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	server = Connection{
		Server:        conn,
	}

	registerClients()

	return &server, nil
}

func registerClients() {
	c := NewGreeterClient(server.Server)
	server.GreeterClient = c
}

func Stop() error {
	return server.Server.Close()
}
