package main

import (
	"errors"
	"fmt"
	"time"

	server "github.com/ummuys/courseproject/grpcway/api"
	handHealth "github.com/ummuys/courseproject/grpcway/handlers/client/health"
	pbHealth "github.com/ummuys/courseproject/grpcway/proto/health/gen"
	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"
	"github.com/ummuys/courseproject/grpcway/tools"
)

func main() {

	//fill db

	fmt.Println("Try to run server . . .")
	serverStatus := make(chan struct{})
	go func() { server.StartListen(serverStatus) }()

	select {
	case <-serverStatus:
		fmt.Println("OK status: server is running\n")
	case <-time.After(1 * time.Second):
		panic(errors.New("can't create a server\n"))
	}

	conn, err := tools.InitClientGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	clientHealth := pbHealth.NewHealthClient(conn)
	clientStudents := pbStudents.NewStudentsClient(conn)
	handHealth.Check(clientHealth)

	tools.Report(1000, clientStudents)
}
