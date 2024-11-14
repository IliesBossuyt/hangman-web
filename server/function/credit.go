package main

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Credit(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("../html/credit.html"))

	// J'execute le template avec les données
	tmpl.Execute(w, nil)
}
