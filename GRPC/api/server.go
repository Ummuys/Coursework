package server

import (
	"net"

	handHealth "coursework/handlers/server/health"
	handStudents "coursework/handlers/server/students"
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
	pbHealth.RegisterHealthServer(grpcServer, &handHealth.Health{})
	pbStudents.RegisterStudentsServer(grpcServer, &handStudents.Students{})
	serverStatus <- struct{}{}

	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}

}
