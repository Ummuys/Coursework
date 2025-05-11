package tools

import (
	"coursework/handlers/client"
	"fmt"
	"slices"
	"time"
)

func Report(n int) {
	times := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		times[i] = client.TestRest()
	}

	fmt.Println("All times\n", times[1:])
	fmt.Println("\nMax time\n", slices.Max(times[1:]))
	fmt.Println("\nMin time\n", slices.Min(times[1:]))

}
