package students

import (
	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"
	"github.com/ummuys/courseproject/grpcway/repository"
)

type Students struct {
	DB repository.DataBase
	pbStudents.UnimplementedStudentsServer
}
