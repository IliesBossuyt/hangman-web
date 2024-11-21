package engine

import (
	"bufio"
	"os"
	"strings"
)

// Fonction pour charger les mots depuis le fichier "words.txt"
func ChargerMotsDepuisFichier() []string {
<<<<<<< Updated upstream:Hangman/loadword.go
	fichier, _ := os.Open("words.txt")
=======
	fichier, _ := os.Open("server/wordlist/words.txt")
>>>>>>> Stashed changes:server/function/loadword.go

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
<<<<<<< Updated upstream:Hangman/loadword.go
	fichier, _ := os.Open("wordshard.txt")
=======
	fichier, _ := os.Open("server/wordlist/wordshard.txt")
>>>>>>> Stashed changes:server/function/loadword.go

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
