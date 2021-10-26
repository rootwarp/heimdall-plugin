package main

import (
	"log"
	"time"

	plugin_grpc "heimdall-plugin/grpc"
	"heimdall-plugin/plugin"
)

const (
	servName = "Node manager heimdall plugin"
)

func main() {
	log.Println("Start ", servName)

	plugin.NewPlugin()
	err := plugin.NewPlugin().Start()
	if err != nil {
		log.Panic(err)
	}

	plugin_grpc.StartServer()

	for {
		// TODO: Handle SIGTERM, Shutdown gracefully.
		time.Sleep(time.Second * 10)
	}
}
