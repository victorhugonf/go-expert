package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.csv").ParseFiles("template.csv"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 100},
		{"Python", 20},
	})
	if err != nil {
		panic(err)
	}
}
