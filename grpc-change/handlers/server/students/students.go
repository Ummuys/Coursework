package students

import (
	pbStudents "github.com/ummuys/coursework/grpc-way/proto/students/gen"
	"github.com/ummuys/coursework/grpc-way/repository"
)

type Students struct {
	DB repository.DataBase
	pbStudents.UnimplementedStudentsServer
}
