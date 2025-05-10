package tools

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
const bletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var types = []string{"Budget", "Paid", "Target"}
var facultys = []string{"AIT", "PGS", "TS", "TES", "UPL", "EIM"}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomType() string {

	return types[rand.Intn(3)]
}

func RandomFaculty() string {
	return facultys[rand.Intn(6)]
}

func RandomGroup() string {
	group := make([]byte, 3)
	for i := range 3 {
		group[i] = bletters[rand.Intn(len(bletters))]
	}
	return string(group)
}
