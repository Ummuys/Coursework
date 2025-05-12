package tools

import (
	"fmt"
	"slices"
	"time"

	"github.com/ummuys/courseproject/restway/handlers/client"
)

func Report(n int) {
	n += 1
	times := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		times[i] = client.TestRest()
	}

	fmt.Println("<====> RESULT REST-WAY <====>\n")
	fmt.Println("Number of tests: ", n-1, "\n")
	fmt.Println("Max time: ", slices.Max(times[1:]), "\n")
	fmt.Println("Min time: ", slices.Min(times[1:]), "\n")
}
