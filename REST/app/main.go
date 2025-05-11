package main

import (
	"fmt"
	"time"

	server "github.com/ummuys/coursework/rest-way/api"
	"github.com/ummuys/coursework/rest-way/tools"
)

func main() {
	fmt.Println("Try to start server . . .")
	go func() { server.InitGin() }()
	fmt.Println("OK status: server is running")

	<-time.After(1000000000 * time.Second)
	tools.Report(1000)
}
