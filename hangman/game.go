
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func (j *Jeu) Jouer() {

	// Initialise les lettres à trouvées à "_"
	for i := range j.LettresaTrouvées {
		if j.LettresaTrouvées[i] == " " {
			j.LettresaTrouvées[i] = " "
		} else {
			j.LettresaTrouvées[i] = "_"
		}
	}

	// Déterminer le nombre de lettres à sélectionner
	nbLettres := rand.Intn(2) + 1

	// Sélectionner les lettres aléatoires dans j.MotADeviner
	indices := make([]int, nbLettres)
	for i := range indices {
		for {
			indices[i] = rand.Intn(len(j.MotADeviner))
			if !contientInt(indices[:i], indices[i]) {
				break
			}
		}
	}

	// Initier les lettres sélectionnées
	lettre := make([]string, nbLettres)
	for i := range indices {
		lettre[i] = string(j.MotADeviner[indices[i]])
	}

	// Ajouter les lettres déjà trouvées
	for i := range j.MotADeviner {
		for _, l := range lettre {
			if j.MotADeviner[i] == l[0] {
				j.LettresaTrouvées[i] = string(j.MotADeviner[i])
			}
		}
	}

	// Ajouter les lettres déjà proposées
	for _, l := range lettre {
		if !contient(j.LettresProposées, l) {
			j.LettresProposées = append(j.LettresProposées, l)
		}
	}

	// Affichage du jeu
	for j.ViesRestantes > 0 && strings.Join(j.LettresaTrouvées, "") != j.MotADeviner {
		fmt.Print("\n")
		afficherEtapePendu(j.ViesRestantes, j.EtapesPendu) // Affiche l'étape du pendu
		fmt.Println("Mot à deviner : ", strings.Join(j.LettresaTrouvées, " "))
		fmt.Println("Il vous reste", j.ViesRestantes, "vies.")
		fmt.Println("Lettres déjà proposées :", j.LettresProposées)
		fmt.Println("Mot déjà proposé :", j.MotProposés)
		fmt.Print("Proposez une lettre ou un mot : ")

		// Saisie de la lettre
		lecteur := bufio.NewReader(os.Stdin)
		saisie, _ := lecteur.ReadString('\n')
		lettre := strings.TrimSpace(strings.ToLower(saisie))

		// Vérifier si la lettre est déjà proposée
		if contient(j.LettresProposées, lettre) {
			fmt.Println("Vous avez déjà proposé cette lettre.")
			continue
		}

		// Vérifier si le mot est déjà proposé
		if contient(j.MotProposés, lettre) {
			fmt.Println("Vous avez déjà proposé ce mot.")
			continue
		}

		// Ajouter la lettre à la liste des lettres trouvées
		if lettre <= "z" && lettre >= "a" && len(lettre) == 1 {
			j.LettresProposées = append(j.LettresProposées, lettre)
		}

		// Ajouter le mot à la liste des mots proposés
		if len(lettre) > 2 {
			j.MotProposés = append(j.MotProposés, lettre)
		}

		// Vérifier si la saisie est égal au mot
		if lettre == j.MotADeviner {
			fmt.Println("Félicitations, vous avez deviné le mot :", j.MotADeviner)
			End()
		} else if len(lettre) > 2 {
			j.ViesRestantes -= 2
			fmt.Println("Mot incorrect.")
		}

		// Vérifier si la lettre est dans le mot
		if strings.Contains(j.MotADeviner, lettre) {
			for i := 0; i < len(j.MotADeviner); i++ {
				if string(j.MotADeviner[i]) == lettre {
					j.LettresaTrouvées[i] = lettre
				}
			}
		} else if lettre <= "z" && lettre >= "a" && len(lettre) == 1 {
			j.ViesRestantes--
			fmt.Println("Lettre incorrecte.")
		}
	}

	// Fin du jeu
	if j.ViesRestantes == 0 {
		fmt.Println("Vous avez perdu. Le mot était :", j.MotADeviner)
		End()
	} else {
		fmt.Println("Félicitations, vous avez deviné le mot :", j.MotADeviner)
		End()
	}
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

// Fonction pour vérifier si un entier est dans un slice
func contientInt(slice []int, item int) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
