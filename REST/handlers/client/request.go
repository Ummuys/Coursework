package client

import (
	"net/http"
	"time"
)

func TestRest() time.Duration {
	start := time.Now()
	info, err := http.Get("http://localhost:8080/students")
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	defer info.Body.Close()
	return end
}
