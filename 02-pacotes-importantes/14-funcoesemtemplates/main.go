package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/csv")
		w.Header().Add("Content-Disposition", "attachment; filename=\"cursos.csv\"")
		templates := []string{
			"title.csv",
			"template.csv",
		}
		t := template.New("template.csv")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
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
