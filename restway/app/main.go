package main

import (
	"fmt"
	"time"

	server "github.com/ummuys/courseproject/restway/api"
	"github.com/ummuys/courseproject/restway/tools"
)

func main() {
	fmt.Println("Try to start server . . .")
	go func() { server.InitGin() }()
	fmt.Println("OK status: server is running\n")

	<-time.After(2 * time.Second)
	tools.Report(1000)
}
