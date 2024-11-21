package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Difficult(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier difficult.html
	tmpl := template.Must(template.ParseFiles("front/template/difficult.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		ViesRestantes: jeu.ViesRestantes,
	}

	// Je redirige vers la page de jeu facile
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "facile" {
			jeu.NouveauJeuFacile()
			jeu.EtapePendu()
			http.Redirect(w, r, "/gameeasy", http.StatusFound)
		}
	}

	// Je redirige vers la page de jeu difficile
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "difficile" {
			jeu.NouveauJeuDifficile()
			jeu.EtapePendu()
			http.Redirect(w, r, "/gamehard", http.StatusFound)
		}
	}

	// Je redirige vers la page de jeu difficile
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "bonus" {
			//jeu.NouveauJeuBonus()
			jeu.EtapePendu()
			http.Redirect(w, r, "/gamebonus", http.StatusFound)
		}
	}

	// Augmenter le nombre de vies de 1
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "button+" && jeu.ViesRestantes < 11 {
			jeu.ViesRestantes++
			w.Header().Set("Refresh", "0")

		}
	}

	// Diminuer le nombre de vies de 1
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "button-" && jeu.ViesRestantes > 1 {
			jeu.ViesRestantes--
			w.Header().Set("Refresh", "0")
		}
	}

	// J'execute le template avec les données
	tmpl.Execute(w, data)

}
