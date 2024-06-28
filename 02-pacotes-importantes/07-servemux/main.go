package main

import "net/http"

func main() {
	finish := make(chan bool)

	mux8080 := http.NewServeMux()
	mux8080.HandleFunc("/", HomeHandler)
	mux8080.Handle("/blog", blog{title: "My blog"})
	go func() {
		http.ListenAndServe(":8080", mux8080)
	}()

	mux8081 := http.NewServeMux()
	mux8081.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Victor!"))
	})
	go func() {
		http.ListenAndServe(":8081", mux8081)
	}()

	<-finish
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
