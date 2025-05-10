package repository

import (
	"coursework/tools"
	"math/rand"
	"strconv"
	"time"
)

var Db = DataBase{}

type StudentInfo struct {
	FirstName     string  `json:"first_name"`
	SecondName    string  `json:"second_name"`
	Surname       string  `json:"surname"`
	Email         string  `json:"email"`
	AverageRating float32 `json:"avg_raiting"`
	Type          string  `json:"type"`
	Faculty       string  `json:"faculty"`
	Group         string  `json:"group"`
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

func (d *DataBase) ToJson() map[string]StudentInfo {
	jsonMap := make(map[string]StudentInfo, len(d.Fields))
	for id, info := range d.Fields {
		key := strconv.FormatInt(id, 10)
		jsonMap[key] = info
	}
	return jsonMap
}
