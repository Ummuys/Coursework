package server

import (
	"net"

	apiHealth "coursework/api/health"
	apiStudents "coursework/api/students"
	pbHealth "coursework/proto/health/gen"
	pbStudents "coursework/proto/students/gen"

	"google.golang.org/grpc"
)

func StartListen(serverStatus chan struct{}) {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pbHealth.RegisterHealthServer(grpcServer, &apiHealth.Health{})
	pbStudents.RegisterStudentsServer(grpcServer, &apiStudents.Students{})
	serverStatus <- struct{}{}

	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}

}
