package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Loose(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/loose.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		MotADeviner:   jeu.MotADeviner,
		Score:         jeu.Score,
		MeilleurScore: jeu.MeilleurScore,
	}

	// Mise à jour du meilleur score
	if  jeu.Score > jeu.MeilleurScore{
		jeu.MeilleurScore = jeu.Score
		w.Header().Set("Refresh", "0")
	}

	// Rediriger vers la séléction de la difficulté pour rejouer
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "Nouvellepartie" {
			// Initialisation du nombre de vies par défaut
			jeu.ViesRestantes = 11
			// Reinitialisation du score
			jeu.Score = 0
			http.Redirect(w, r, "/difficult", http.StatusFound)
		}
	}

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}
