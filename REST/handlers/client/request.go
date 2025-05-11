package client

import (
	"fmt"
	"net/http"
	"time"
)

func TestRest() {
	start := time.Now()
	info, err := http.Get("http://localhost:8080/students")
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	defer info.Body.Close()
	fmt.Println(info.Body)
	fmt.Println("Time taken --> ", end)
}
