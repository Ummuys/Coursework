package health

import pbHealth "github.com/ummuys/courseproject/grpcway/proto/health/gen"

type Health struct {
	pbHealth.UnimplementedHealthServer
}
