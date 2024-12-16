package engine

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func (jeu *Engine) GameBonus(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("front/template/gamebonus.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		MotADeviner:      strings.Join(jeu.LettresaTrouvées, " "),
		ViesRestantes:    jeu.ViesRestantes,
		LettresProposées: jeu.LettresProposées,
		MotProposés:      jeu.MotProposés,
		Message:          jeu.Message,
		EtapesPendu:      jeu.EtapesPendu,
	}

	// Réinitialiser jeu.Message
	jeu.Message = ""

	// Récupèrer le mot ou la lettre séléctionné par le joueur
	if r.Method == "POST" {
		mot := r.FormValue("mot")

		// Vérifier si la lettre est déjà proposée
		if contient(jeu.LettresProposées, mot) {
			jeu.Message = ("Vous avez déjà proposé cette lettre.")
			http.Redirect(w, r, "/gamebonus", http.StatusFound)
			return
		}

		// Vérifier si le mot est déjà proposé
		if contient(jeu.MotProposés, mot) {
			jeu.Message = ("Vous avez déjà proposé ce mot.")
			http.Redirect(w, r, "/gamebonus", http.StatusFound)
			return
		}

		// Vérifier si la lettre est dans le mot
		if strings.Contains(jeu.MotADeviner, mot) {
			for i := 0; i < len(jeu.MotADeviner); i++ {
				if string(jeu.MotADeviner[i]) == mot {
					jeu.LettresaTrouvées[i] = mot
					jeu.Message = ("Bonne lettre !")
					// Jouer le son
					fmt.Fprintf(w, `
					<script>
						window.onload = function() {
							var audio = new Audio('/static/song/correct.mp3');
							audio.play();
							audio.onended = function() {
								location.replace(location.href);
							};
						};
					</script>`)
				}
			}
		} else if mot <= "z" && mot >= "a" && len(mot) == 1 {
			jeu.Message = ("Mauvaise lettre !")
			jeu.ViesRestantes--
			// Avancer le pendu
			jeu.EtapePendu()
			// Jouer le son
			fmt.Fprintf(w, `
			<script>
				window.onload = function() {
					var audio = new Audio('/static/song/wrong.mp3');
					audio.play();
					audio.onended = function() {
						location.replace(location.href);
					};
				};
			</script>`)

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
			// Avancer le pendu
			jeu.EtapePendu()
			// Jouer le son
			fmt.Fprintf(w, `
			<script>
				window.onload = function() {
					var audio = new Audio('/static/song/wrong.mp3');
					audio.play();
					audio.onended = function() {
						location.replace(location.href);
					};
				};
			</script>`)
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
