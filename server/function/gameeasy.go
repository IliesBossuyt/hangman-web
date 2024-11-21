package engine

import (
	"html/template"
	"net/http"
	"strings"
)

func (jeu *Engine) GameEasy(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/gameeasy.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		MotADeviner:      strings.Join(jeu.LettresaTrouvées, " "),
		ViesRestantes:    jeu.ViesRestantes,
		LettresProposées: jeu.LettresProposées,
		MotProposés:      jeu.MotProposés,
		Message:          jeu.Message,
		EtapesPendu:      jeu.EtapesPendu,
	}

	jeu.Message = ""

	if r.Method == "POST" {
		mot := r.FormValue("mot")

		// Vérifier si la lettre est déjà proposée
		if contient(jeu.LettresProposées, mot) {
			jeu.Message = ("Vous avez déjà proposé cette lettre.")
			http.Redirect(w, r, "/gameeasy", http.StatusFound)
			return
		}

		// Vérifier si le mot est déjà proposé
		if contient(jeu.MotProposés, mot) {
			jeu.Message = ("Vous avez déjà proposé ce mot.")
			http.Redirect(w, r, "/gameeasy", http.StatusFound)
			return
		}

		// Vérifier si la lettre est dans le mot
		if strings.Contains(jeu.MotADeviner, mot) {
			for i := 0; i < len(jeu.MotADeviner); i++ {
				if string(jeu.MotADeviner[i]) == mot {
					jeu.LettresaTrouvées[i] = mot
					jeu.Message = ("Bonne lettre !")
					w.Header().Set("Refresh", "0")
				}
			}
		} else if mot <= "z" && mot >= "a" && len(mot) == 1 {
			jeu.Message = ("Mauvaise lettre !")
			jeu.ViesRestantes--
			jeu.EtapePendu()
			w.Header().Set("Refresh", "0")
		}

		// Ajouter la lettre à la liste des lettres proposées
		if mot <= "z" && mot >= "a" && len(mot) == 1 {
			jeu.LettresProposées = append(jeu.LettresProposées, mot)
		}

		// Ajouter le mot à la liste des mots proposés
		if len(mot) > 2 {
			jeu.MotProposés = append(jeu.MotProposés, mot)
		}

		// Vérifier si la saisie est égal au mot
		if mot == jeu.MotADeviner {
			http.Redirect(w, r, "/win", http.StatusFound)
		} else if len(mot) > 2 {
			jeu.ViesRestantes -= 2
			jeu.Message = ("Mot incorrect !")
			jeu.EtapePendu()
			w.Header().Set("Refresh", "0")
		}

	}

	// Fin du jeu
	if jeu.ViesRestantes <= 0 {
		http.Redirect(w, r, "/loose", http.StatusFound)
	} else if strings.Join(jeu.LettresaTrouvées, "") == jeu.MotADeviner {
		http.Redirect(w, r, "/win", http.StatusFound)
	}
	// J'execute le template avec les données
	tmpl.Execute(w, data)

}

// Fonction pour vérifier si un mot est dans un slice
func contient(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
