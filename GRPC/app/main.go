package main

import (
	server "coursework/api"
	handHealth "coursework/handlers/client/health"
	handStudents "coursework/handlers/client/students"
	pbHealth "coursework/proto/health/gen"
	pbStudents "coursework/proto/students/gen"
	"coursework/repository"
	"coursework/tools"
	"errors"
	"fmt"
	"time"
)

func main() {

	//fill db
	fmt.Println("Try to fill data. . .")
	repository.Db.SetFields(1000)
	fmt.Println("OK status: data set")

	fmt.Println("Try to run server . . .")
	serverStatus := make(chan struct{})
	go func() { server.StartListen(serverStatus) }()

	select {
	case <-serverStatus:
		fmt.Println("OK status: server is running")
	case <-time.After(1 * time.Second):
		panic(errors.New("can't create a server"))
	}

	conn, err := tools.InitGRPC()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Println("Conn is succsesful")
	clientHealth := pbHealth.NewHealthClient(conn)
	clientStudents := pbStudents.NewStudentsClient(conn)
	handHealth.Check(clientHealth)
	time.Sleep(3 * time.Second)
	handStudents.Get(clientStudents)
}
