package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", jeu.Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	http.HandleFunc("/difficult", jeu.Difficult)
	http.HandleFunc("/gameeasy", jeu.GameEasy)
	http.HandleFunc("/gamehard", jeu.GameHard)
	http.HandleFunc("/pause", jeu.Pause)
	http.HandleFunc("/credit", jeu.Credit)
	http.HandleFunc("/win", jeu.Win)
	http.HandleFunc("/loose", jeu.Loose)
	http.HandleFunc("/gamebonus", jeu.GameBonus)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hangman", jeu.Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	// On lance le serveur local sur le port 8080
}
