package main

import (
	"aprendendo-go/exercicios"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	MapsOutput    string
	StructsOutput string
}

var data PageData

func main() {
	// Carrega o template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Rota principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	// Rota para rodar Maps
	http.HandleFunc("/run/maps", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data.MapsOutput = exercicios.ExercicioMaps()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Rota para rodar Structs
	http.HandleFunc("/run/structs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data.StructsOutput = exercicios.ExercicioStructs()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("🚀 Dashboard rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
