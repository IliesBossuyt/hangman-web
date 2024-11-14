package main

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Pause(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("../html/pause.html"))

	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "Nouvellepartie" {
			jeu.ViesRestantes = 11
			http.Redirect(w, r, "/difficult", http.StatusFound)
		}
	}

	// J'execute le template avec les données
	tmpl.Execute(w, nil)
}
