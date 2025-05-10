package students

import (
	"context"
	pbStudents "coursework/proto/students/gen"
	repos "coursework/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Students) Get(ctx context.Context, _ *emptypb.Empty) (*pbStudents.GetResponse, error) {
	return &pbStudents.GetResponse{
		Students: repos.Db.ToProto(),
	}, nil
}
