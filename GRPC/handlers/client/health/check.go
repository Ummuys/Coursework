package health

import (
	"context"
	pbHealth "coursework/proto/health/gen"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

func Check(client pbHealth.HealthClient) {
	start := time.Now()
	_, err := client.Check(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	end := time.Since(start)
	fmt.Println("Time taken -> ", end)
}
