package main

import (
	"errors"
	"fmt"
	"time"

	server "github.com/ummuys/coursework/grpc-way/api"
	handHealth "github.com/ummuys/coursework/grpc-way/handlers/client/health"
	pbHealth "github.com/ummuys/coursework/grpc-way/proto/health/gen"
	pbStudents "github.com/ummuys/coursework/grpc-way/proto/students/gen"
	"github.com/ummuys/coursework/grpc-way/tools"
)

func main() {

	//fill db

	fmt.Println("Try to run server . . .")
	serverStatus := make(chan struct{})
	go func() { server.StartListen(serverStatus) }()

	select {
	case <-serverStatus:
		fmt.Println("OK status: server is running")
	case <-time.After(1 * time.Second):
		panic(errors.New("can't create a server"))
	}

	conn, err := tools.InitClientGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Conn is succsesful")

	clientHealth := pbHealth.NewHealthClient(conn)
	clientStudents := pbStudents.NewStudentsClient(conn)
	handHealth.Check(clientHealth)

	tools.Report(1000, clientStudents)
}
