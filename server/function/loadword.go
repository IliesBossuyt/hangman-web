package engine

import (
	"bufio"
	"os"
	"strings"
)

// Fonction pour charger les mots depuis le fichier "words.txt"
func ChargerMotsDepuisFichier() []string {
	fichier, _ := os.Open("server/wordlist/words.txt")

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

// Fonction pour charger les mots depuis le fichier "wordshard.txt"
func ChargerMotsDepuisFichierHard() []string {
	fichier, _ := os.Open("server/wordlist/wordshard.txt")

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
