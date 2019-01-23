package main

import (
	"fmt"

	"github.com/michalnov/SovyGo/bin/server"
)

func main() {
	fmt.Println("Hello Server")

	running := make(chan int)

	go func() {
		server := server.Server{}
		server.SetupServer(running)
		server.StartServer()
	}()

	for i := 0; i < 1; i++ {
		_ = <-running
	}
}
