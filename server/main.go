package main

import (
	"bufio"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Je crée ma structure
type Engine struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
	EtapesPendu      string
	Message          string
}

func main() {
	var jeu Engine
	http.HandleFunc("/", jeu.Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	http.HandleFunc("/difficult", jeu.Difficult)
	http.HandleFunc("/gameeasy", jeu.GameEasy)
	http.HandleFunc("/gamehard", jeu.GameHard)
	http.HandleFunc("/pause", jeu.Pause)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hangman", jeu.Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	// On lance le serveur local sur le port 8080
}

func (jeu *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("start.html"))

	// Définir le nombre de vies par défaut a 11
	jeu.ViesRestantes = 11

	// J'execute le template avec les données
	tmpl.Execute(w, nil)
}

func (jeu *Engine) Pause(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("pause.html"))

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

func (jeu *Engine) Difficult(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier difficult.html
	tmpl := template.Must(template.ParseFiles("difficult.html"))

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

func (jeu *Engine) GameEasy(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("gameeasy.html"))

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
			jeu.Message = ("Félicitations, vous avez deviné le mot : " + jeu.MotADeviner)
			w.Header().Set("Refresh", "0")
			//End()
		} else if len(mot) > 2 {
			jeu.ViesRestantes -= 2
			jeu.Message = ("Mot incorrect !")
			jeu.EtapePendu()
			w.Header().Set("Refresh", "0")
		}

	}

	// Fin du jeu
	if jeu.ViesRestantes == 0 {
		jeu.Message = ("Vous avez perdu. Le mot était : " + jeu.MotADeviner)
		//End()
	} else if strings.Join(jeu.LettresaTrouvées, "") == jeu.MotADeviner {
		jeu.Message = ("Félicitations, vous avez deviné le mot : " + jeu.MotADeviner)
		//End()
	}
	// J'execute le template avec les données
	tmpl.Execute(w, data)

}

func (jeu *Engine) GameHard(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("gamehard.html"))

	// Je crée une variable qui définit ma structure
	data := Engine{
		MotADeviner:      strings.Join(jeu.LettresaTrouvées, " "),
		ViesRestantes:    jeu.ViesRestantes,
		LettresProposées: jeu.LettresProposées,
		MotProposés:      jeu.MotProposés,
		Message:          jeu.Message,
	}

	jeu.Message = ""

	if r.Method == "POST" {
		mot := r.FormValue("mot")

		// Vérifier si la lettre est déjà proposée
		if contient(jeu.LettresProposées, mot) {
			jeu.Message = ("Vous avez déjà proposé cette lettre.")
			http.Redirect(w, r, "/gamehard", http.StatusFound)
			return
		}

		// Vérifier si le mot est déjà proposé
		if contient(jeu.MotProposés, mot) {
			jeu.Message = ("Vous avez déjà proposé ce mot.")
			http.Redirect(w, r, "/gamehard", http.StatusFound)
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
			jeu.Message = ("Félicitations, vous avez deviné le mot : " + jeu.MotADeviner)
			w.Header().Set("Refresh", "0")
			//End()
		} else if len(mot) > 2 {
			jeu.ViesRestantes -= 2
			jeu.Message = ("Mot incorrect !")
			jeu.EtapePendu()
			w.Header().Set("Refresh", "0")
		}

	}

	// Fin du jeu
	if jeu.ViesRestantes == 0 {
		jeu.Message = ("Vous avez perdu. Le mot était : " + jeu.MotADeviner)
		//End()
	} else if strings.Join(jeu.LettresaTrouvées, "") == jeu.MotADeviner {
		jeu.Message = ("Félicitations, vous avez deviné le mot : " + jeu.MotADeviner)
		//End()
	}
	// J'execute le template avec les données
	tmpl.Execute(w, data)

}

func (jeu *Engine) EtapePendu() {
	if jeu.ViesRestantes == 11 {
		jeu.EtapesPendu = "/static/1.png"
	}
	if jeu.ViesRestantes == 10 {
		jeu.EtapesPendu = "/static/2.png"
	}
	if jeu.ViesRestantes == 9 {
		jeu.EtapesPendu = "/static/3.png"
	}
	if jeu.ViesRestantes == 8 {
		jeu.EtapesPendu = "/static/4.png"
	}
	if jeu.ViesRestantes == 7 {
		jeu.EtapesPendu = "/static/5.png"
	}
	if jeu.ViesRestantes == 6 {
		jeu.EtapesPendu = "/static/6.png"
	}
	if jeu.ViesRestantes == 5 {
		jeu.EtapesPendu = "/static/7.png"
	}
	if jeu.ViesRestantes == 4 {
		jeu.EtapesPendu = "/static/8.png"
	}
	if jeu.ViesRestantes == 3 {
		jeu.EtapesPendu = "/static/9.png"
	}
	if jeu.ViesRestantes == 2 {
		jeu.EtapesPendu = "/static/10.png"
	}
	if jeu.ViesRestantes == 1 {
		jeu.EtapesPendu = "/static/11.png"
	}
}

var mots []string

func (jeu *Engine) NouveauJeuFacile() {
	// Charger les mots
	mots = ChargerMotsDepuisFichier()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	jeu.MotADeviner = motSansAccents
	jeu.LettresaTrouvées = make([]string, len(motSansAccents))
	jeu.LettresProposées = []string{}
	jeu.MotProposés = []string{}

	// Initialise les lettres à trouvées à "_"
	for i := range jeu.LettresaTrouvées {
		if jeu.LettresaTrouvées[i] == " " {
			jeu.LettresaTrouvées[i] = " "
		} else {
			jeu.LettresaTrouvées[i] = "_"
		}
	}

}

// Fonction pour créer le jeu en mode difficile
func (jeu *Engine) NouveauJeuDifficile() {
	// Charger les mots
	mots = ChargerMotsDepuisFichierHard()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	jeu.MotADeviner = motSansAccents
	jeu.LettresaTrouvées = make([]string, len(motSansAccents))
	jeu.LettresProposées = []string{}
	jeu.MotProposés = []string{}

	// Initialise les lettres à trouvées à "_"
	for i := range jeu.LettresaTrouvées {
		if jeu.LettresaTrouvées[i] == " " {
			jeu.LettresaTrouvées[i] = " "
		} else {
			jeu.LettresaTrouvées[i] = "_"
		}
	}
}

// Fonction pour enlever les accents
func enleverAccents(mot string) string {
	accents := []string{"é", "è", "ê", "ë", "ï", "î", "ô", "ö", "ù", "ü", "û", "à", "â", "ä", "ç", "ÿ", "œ", "æ", "ᵫ", "ꭣ", "ꭡ"}
	nonAccents := []string{"e", "e", "e", "e", "i", "i", "o", "o", "u", "u", "u", "a", "a", "a", "c", "y", "oe", "ae", "ue", "uo", "ie"}
	for i, accent := range accents {
		mot = strings.Replace(mot, accent, nonAccents[i], -1)
	}
	return mot
}

func enleverMajuscules(mot string) string {
	majuscules := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "Á", "À", "Â", "Ã", "Å", "Ä", "Ç", "È", "É", "Ê", "Ë", "Î", "Ì", "Ï", "Ñ", "Ô", "Õ", "Û", "Ù", "Ü", "Ÿ", "Æ", "Ꜵ", "Œ"}
	nonMajuscules := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "a", "a", "a", "a", "a", "a", "c", "e", "e", "e", "e", "i", "i", "i", "n", "o", "o", "u", "u", "u", "y", "ae", "ao", "oe"}
	for i, accent := range majuscules {
		mot = strings.Replace(mot, accent, nonMajuscules[i], -1)
	}
	return mot
}

// Fonction pour charger les mots depuis le fichier "words.txt"
func ChargerMotsDepuisFichier() []string {
	fichier, _ := os.Open("words.txt")

	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)

	// Lire chaque ligne et ajouter à la liste des mots
	for scanner.Scan() {
		mot := strings.TrimSpace(scanner.Text())
		if mot != "" {
			mots = append(mots, mot)
		}
	}

	return mots
}

func ChargerMotsDepuisFichierHard() []string {
	fichier, _ := os.Open("wordshard.txt")

	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)

	// Lire chaque ligne et ajouter à la liste des mots
	for scanner.Scan() {
		mot := strings.TrimSpace(scanner.Text())
		if mot != "" {
			mots = append(mots, mot)
		}
	}

	return mots
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
