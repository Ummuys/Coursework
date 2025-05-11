package tools

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClientGRPC() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
