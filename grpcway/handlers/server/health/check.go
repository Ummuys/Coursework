package health

import (
	"context"
	"net/http"

	pbHealth "github.com/ummuys/courseproject/grpcway/proto/health/gen"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Health) Check(ctx context.Context, _ *emptypb.Empty) (*pbHealth.CheckResponse, error) {
	return &pbHealth.CheckResponse{Status: http.StatusOK}, nil
}
