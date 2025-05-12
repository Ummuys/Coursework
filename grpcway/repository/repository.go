package repository

import (
	"math/rand"
	"time"

	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"
	"github.com/ummuys/courseproject/grpcway/tools"
)

type DataBase interface {
	SetFields(n int)
	ToProto()
	GetBasicFields() map[int64]StudentInfo
	GetProtoFields() map[int64]*pbStudents.StudentInfo
}

type StudentInfo struct {
	FirstName     string
	SecondName    string
	Surname       string
	Email         string
	AverageRating float32
	Type          string
	Faculty       string
	Group         string
}

type MapDb struct {
	BasicFields map[int64]StudentInfo
	ProtoFields map[int64]*pbStudents.StudentInfo
}

func NewMapDB() DataBase {
	return &MapDb{
		BasicFields: make(map[int64]StudentInfo),
		ProtoFields: make(map[int64]*pbStudents.StudentInfo),
	}
}

func (db *MapDb) SetFields(n int) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i <= n; i++ {
		db.BasicFields[int64(i)] = StudentInfo{
			FirstName:     tools.RandomString(6),
			SecondName:    tools.RandomString(6),
			Surname:       tools.RandomString(6),
			Email:         tools.RandomString(8) + "@gmail.com",
			AverageRating: 2.00 + rand.Float32()*(5.01-2.00),
			Type:          tools.RandomType(),
			Faculty:       tools.RandomFaculty(),
			Group:         tools.RandomGroup(),
		}
	}
}

func (d *MapDb) GetBasicFields() map[int64]StudentInfo {
	return d.BasicFields
}

func (d *MapDb) GetProtoFields() map[int64]*pbStudents.StudentInfo {
	return d.ProtoFields
}

func (d *MapDb) ToProto() {
	for id, r := range d.BasicFields {
		d.ProtoFields[id] = &pbStudents.StudentInfo{
			FirstName:     r.FirstName,
			SecondName:    r.SecondName,
			Surname:       r.Surname,
			Email:         r.Email,
			AverageRating: r.AverageRating,
			Type:          r.Type,
			Faculty:       r.Faculty,
			Group:         r.Group,
		}
	}
}
