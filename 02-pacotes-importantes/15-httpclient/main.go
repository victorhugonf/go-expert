package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Microsecond}
	request, err := c.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(response))

}
