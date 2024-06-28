package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	json := bytes.NewBuffer([]byte(`{"name":"victor"}`))
	request, err := c.Post("https://www.google.com", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	io.CopyBuffer(os.Stdout, request.Body, nil)
}
