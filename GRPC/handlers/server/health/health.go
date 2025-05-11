package health

import pbHealth "github.com/ummuys/coursework/grpc-way/proto/health/gen"

type Health struct {
	pbHealth.UnimplementedHealthServer
}
