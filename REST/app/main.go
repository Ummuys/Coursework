package main

import (
	repos "coursework/repository"
	client "coursework/web/api/client"
	"coursework/web/server"
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
