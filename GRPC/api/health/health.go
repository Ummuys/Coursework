package health

import pbHealth "coursework/proto/health/gen"

type Health struct {
	pbHealth.UnimplementedHealthServer
}
