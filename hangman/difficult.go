
package main

import (
	"math/rand"
	"os"
	"time"
)

var mots []string

type Jeu struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
	EtapesPendu      []string
}

// Fonction pour créer le jeu en mode facile
func NouveauJeuFacile() *Jeu {
	// Charger les mots selon la difificulté des mots choisits
	mode := os.Args[1]
	if mode == "hard" {
		mots = ChargerMotsDepuisFichierHard()
	} else if mode == "normal" {
		mots = ChargerMotsDepuisFichier()
	}

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
	// Charger les mots selon la difificulté des mots choisits
	mode := os.Args[1]
	if mode == "hard" {
		mots = ChargerMotsDepuisFichierHard()
	} else if mode == "normal" {
		mots = ChargerMotsDepuisFichier()
	}

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
