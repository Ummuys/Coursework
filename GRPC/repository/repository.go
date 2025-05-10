package repository

import (
	pbStudents "coursework/proto/students/gen"
	"coursework/tools"
	"math/rand"
	"time"
)

var Db = DataBase{}

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

type DataBase struct {
	Fields map[int64]StudentInfo
}

func (d *DataBase) SetFields(n int) {
	d.Fields = make(map[int64]StudentInfo, n)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i <= n; i++ {
		d.Fields[int64(i)] = StudentInfo{
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

func (d *DataBase) GetData() map[int64]StudentInfo {
	return d.Fields
}

func (d *DataBase) ToProto() map[int64]*pbStudents.StudentInfo {
	protoMap := make(map[int64]*pbStudents.StudentInfo, len(d.Fields))

	for id, r := range d.Fields {
		protoMap[id] = &pbStudents.StudentInfo{
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
	return protoMap
}
