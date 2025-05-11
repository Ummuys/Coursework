package students

import pbStudents "coursework/proto/students/gen"

type Students struct {
	pbStudents.UnimplementedStudentsServer
}
