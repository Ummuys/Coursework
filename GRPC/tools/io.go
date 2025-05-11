package tools

import (
	"fmt"
	"slices"
	"time"

	hand "github.com/ummuys/coursework/grpc-way/handlers/client/students"
	pbStudents "github.com/ummuys/coursework/grpc-way/proto/students/gen"
)

func Report(n int, clientStudents pbStudents.StudentsClient) {
	times := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		times[i] = hand.Get(clientStudents)
	}

	fmt.Println("All times\n", times)
	fmt.Println("\nMax time\n", slices.Max(times))
	fmt.Println("\nMin time\n", slices.Min(times))
}
