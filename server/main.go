package main

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Je crée ma structure
type Jeu struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
	EtapesPendu      []string
}

func main() {
	http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	http.HandleFunc("/difficult", Difficult)
	http.HandleFunc("/gameeasy", GameEasy)
	http.HandleFunc("/gamehard", GameHard)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hangman", Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	// On lance le serveur local sur le port 8080
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Récupérer le mot soumis
		mot := r.FormValue("mot")

		fmt.Println("Le mot soumis est :", mot)

		// Répondre avec un JSON pour indiquer que le mot a été soumis avec succès
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": true}`))
	} else {
		// Générer un mot aléatoirement
		mot := NouveauJeuFacile().MotADeviner

		// Stocker le mot dans un cookie
		cookie := &http.Cookie{
			Name:  "mot",
			Value: mot,
			Path:  "/",
		}
		http.SetCookie(w, cookie)

		// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
		tmpl := template.Must(template.ParseFiles("start.html"))
		// J'execute le template avec les données
		tmpl.Execute(w, nil)
	}
}

func Difficult(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier difficult.html
	tmpl := template.Must(template.ParseFiles("difficult.html"))
	// J'execute le template avec les données
	tmpl.Execute(w, nil)

}

func GameEasy(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("gameeasy.html"))
	// Je crée une variable qui définit ma structure
	data := Jeu{
		MotADeviner: NouveauJeuFacile().MotADeviner,
	}
	if r.Method == "POST" {
		mot := r.FormValue("mot")
		fmt.Println("Le mot entré est :", mot)
	}
	// J'execute le template avec les données
	tmpl.Execute(w, data)

}

func GameHard(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("gamehard.html"))
	// Je crée une variable qui définit ma structure
	data := Jeu{
		MotADeviner: NouveauJeuDifficile().MotADeviner,
	}
	// J'execute le template avec les données
	tmpl.Execute(w, data)

}

var mots []string

func NouveauJeuFacile() *Jeu {
	// Charger les mots
	mots = ChargerMotsDepuisFichier()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	// Charger les étapes du pendu depuis le fichier hangman.txt
	etapes := chargerEtapesPendu("hangman.txt")

	return &Jeu{
		MotADeviner:      motSansAccents,
		LettresaTrouvées: make([]string, len(motSansAccents)),
		ViesRestantes:    10,
		LettresProposées: []string{},
		MotProposés:      []string{},
		EtapesPendu:      etapes, // Associer les étapes du pendu
	}
}

// Fonction pour créer le jeu en mode difficile
func NouveauJeuDifficile() *Jeu {
	// Charger les mots
	mots = ChargerMotsDepuisFichierHard()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motAleatoire)

	// Charger les étapes du pendu depuis le fichier hangman.txt
	etapes := chargerEtapesPendu("hangman.txt")

	return &Jeu{
		MotADeviner:      motSansAccents,
		LettresaTrouvées: make([]string, len(motSansAccents)),
		ViesRestantes:    5,
		LettresProposées: []string{},
		MotProposés:      []string{},
		EtapesPendu:      etapes, // Associer les étapes du pendu
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

// Fonction pour charger les étapes du pendu depuis le fichier hangman.txt
func chargerEtapesPendu(cheminFichier string) []string {
	fichier, _ := os.Open(cheminFichier)
	defer fichier.Close()

	var etapes []string
	var etape strings.Builder

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			etapes = append(etapes, etape.String())
			etape.Reset()
		} else {
			etape.WriteString(ligne + "\n")
		}
	}

	if etape.Len() > 0 {
		etapes = append(etapes, etape.String())
	}

	return etapes
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
