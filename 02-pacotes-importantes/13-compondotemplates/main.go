package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/csv")
		w.Header().Add("Content-Disposition", "attachment; filename=\"cursos.csv\"")
		templates := []string{
			"title.csv",
			"template.csv",
		}
		t := template.Must(template.New("template.csv").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 100},
			{"Python", 20},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
