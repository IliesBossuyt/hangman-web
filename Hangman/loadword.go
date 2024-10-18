package main

import (
	"bufio"
	"os"
	"strings"
)

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
