package tools

import (
	"fmt"
	"slices"
	"time"

	hand "github.com/ummuys/courseproject/grpcway/handlers/client/students"
	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"
)

func Report(n int, clientStudents pbStudents.StudentsClient) {
	times := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		times[i] = hand.Get(clientStudents)
	}

	fmt.Println("<====> RESULT GRPC-WAY <====>\n")
	fmt.Println("Number of tests: ", n, "\n")
	fmt.Println("Max time: ", slices.Max(times), "\n")
	fmt.Println("Min time: ", slices.Min(times), "\n")
}
