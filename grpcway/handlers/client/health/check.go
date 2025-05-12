package health

import (
	"context"

	pbHealth "github.com/ummuys/courseproject/grpcway/proto/health/gen"

	"google.golang.org/protobuf/types/known/emptypb"
)

func Check(client pbHealth.HealthClient) {
	_, err := client.Check(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
}
