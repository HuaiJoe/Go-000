package main

import (
	"tcp/pkg/server"
)

func main() {
	server := server.NewTCPServer("localhost", 8080)
	server.Start()
}
