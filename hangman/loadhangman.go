package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func afficherEtapePendu(viesRestantes int, etapes []string) {
	erreurs := len(etapes) - viesRestantes
	if erreurs >= len(etapes) {
		fmt.Println("Le joueur a perdu !")
		return
	}
	fmt.Println(etapes[erreurs])
}
