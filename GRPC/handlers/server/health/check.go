package health

import (
	"context"
	pbHealth "coursework/proto/health/gen"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Health) Check(ctx context.Context, _ *emptypb.Empty) (*pbHealth.CheckResponse, error) {
	return &pbHealth.CheckResponse{Status: http.StatusOK}, nil
}
