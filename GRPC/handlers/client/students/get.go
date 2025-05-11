package students

import (
	"context"
	"time"

	pbStudents "github.com/ummuys/coursework/grpc-way/proto/students/gen"

	"google.golang.org/protobuf/types/known/emptypb"
)

func Get(client pbStudents.StudentsClient) time.Duration {
	start := time.Now()
	_, err := client.Get(context.Background(), &emptypb.Empty{})
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	return end
}
