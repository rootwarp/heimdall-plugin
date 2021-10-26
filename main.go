package main

import (
	"log"
	"time"

	plugin_grpc "heimdall-plugin/grpc"
)

const (
	servName = "Node manager heimdall plugin"
)

func main() {
	log.Println("Start ", servName)

	plugin_grpc.StartServer()

	for {
		// TODO: Handle SIGTERM
		time.Sleep(time.Second * 10)
	}
}
