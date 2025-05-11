package students

import (
	"context"
	pbStudents "coursework/proto/students/gen"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

func Get(client pbStudents.StudentsClient) {
	start := time.Now()
	_, err := client.Get(context.Background(), &emptypb.Empty{})
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	fmt.Println("Time taken -> ", end)
}
