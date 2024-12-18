package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/start.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{}

	// Définir le nombre de vies par défaut a 11
	jeu.ViesRestantes = 11

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}
