package repository

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ummuys/courseproject/restway/tools"
)

type DataBase interface {
	SetFields(n int)
	ToJSON()
	GetBasicFields() map[int64]StudentInfo
	GetJSONFields() map[string]StudentInfo
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
	JSONFields  map[string]StudentInfo
}

func NewMapDB() DataBase {
	return &MapDb{
		BasicFields: make(map[int64]StudentInfo),
		JSONFields:  make(map[string]StudentInfo),
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

func (db *MapDb) GetBasicFields() map[int64]StudentInfo {
	return db.BasicFields
}

func (db *MapDb) GetJSONFields() map[string]StudentInfo {
	return db.JSONFields
}

func (db *MapDb) ToJSON() {
	for id, info := range db.BasicFields {
		key := strconv.FormatInt(id, 10)
		db.JSONFields[key] = info
	}
}
