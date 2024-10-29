package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CommencerJeu() {
	for {

		// Vérifier les arguments de la ligne de commande pour sélectionner la liste de mots
		if len(os.Args) > 1 {
			mode := os.Args[1]
			if mode == "hard" {
				ChargerMotsDepuisFichierHard()
			} else if mode == "normal" {
				ChargerMotsDepuisFichier()
			} else {
				fmt.Println("Argument invalide. Utilisez 'hard' ou 'normal'.")
				return
			}
		} else {
			fmt.Println("Veuillez fournir un argument : 'hard' ou 'normal'.")
			return
		}

		fmt.Println(lettreAsciiArt())
		fmt.Println("Veuillez sélectionner le niveau facile ou difficile.")

		// Saisie du mot
		lecteur := bufio.NewReader(os.Stdin)
		saisie, _ := lecteur.ReadString('\n')
		lettre := strings.TrimSpace(strings.ToLower(saisie))

		// Mode de jeu facile
		if lettre == "Facile" || lettre == "facile" {
			jeu := NouveauJeuFacile()
			jeu.Jouer()
			break
		}
		// Mode de jeu difficile
		if lettre == "Difficile" || lettre == "difficile" {
			jeu := NouveauJeuDifficile()
			jeu.Jouer()
			break
		} else {
			fmt.Println("Entrez facile ou difficile")
		}
	}
}
