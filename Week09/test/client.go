package main

import "tcp/pkg/client"

func main() {
	client.Client("localhost", 8080)
}
