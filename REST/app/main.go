package main

import (
	server "coursework/api"
	client "coursework/handlers/client"
	repos "coursework/repository"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Try to start server . . .")
	go func() { server.InitGin() }()
	fmt.Println("OK status: server is running")

	<-time.After(1 * time.Second)

	repos.Db.SetFields(1000)

	client.TestRest()
}
