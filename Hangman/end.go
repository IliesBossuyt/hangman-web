package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func End() {
	fmt.Println("Voulez vous rejouer ?")

	// Saisie du mot
	lecteur := bufio.NewReader(os.Stdin)
	saisie, _ := lecteur.ReadString('\n')
	lettre := strings.TrimSpace(strings.ToLower(saisie))

	// Si non, fin du jeu, Si oui, rejouer
	if lettre == "Non" || lettre == "non" {
		os.Exit(0)
	} else if lettre == "Oui" || lettre == "oui" {
		CommencerJeu()
	} else {
		fmt.Println("Entrez Oui ou Non")
	}
	End()
}
