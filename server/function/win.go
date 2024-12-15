package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Win(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/win.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		MotADeviner:   jeu.MotADeviner,
		ViesRestantes: jeu.ViesRestantes,
		Score:         jeu.Score,
	}

	// Rediriger vers la séléction de la difficulté pour rejouer
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "Nouvellepartie" {
			// Initialisation du nombre de vies par défaut
			jeu.ViesRestantes = 11
			http.Redirect(w, r, "/difficult", http.StatusFound)
		}
	}

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}
