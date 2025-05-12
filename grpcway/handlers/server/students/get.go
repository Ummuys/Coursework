package students

import (
	"context"

	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Students) Get(ctx context.Context, _ *emptypb.Empty) (*pbStudents.GetResponse, error) {
	return &pbStudents.GetResponse{
		Students: s.DB.GetProtoFields(),
	}, nil
}
