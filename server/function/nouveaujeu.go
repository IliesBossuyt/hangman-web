package engine

import (
	"math/rand"
	"time"
)

var mots []string

func (jeu *Engine) NouveauJeuFacile() {
	// Charger les mots
	mots = ChargerMotsDepuisFichier()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	jeu.MotADeviner = motSansAccents
	jeu.LettresaTrouvées = make([]string, 0)
	for _, r := range jeu.MotADeviner {
		jeu.LettresaTrouvées = append(jeu.LettresaTrouvées, string(r))
	}
	jeu.LettresProposées = []string{}
	jeu.MotProposés = []string{}

	// Initialise les lettres à trouvées à "_"
	for i := 0; i < len(jeu.LettresaTrouvées); i++ {
		if jeu.LettresaTrouvées[i] == " " {
			jeu.LettresaTrouvées[i] = " "
		} else {
			jeu.LettresaTrouvées[i] = "_"
		}
	}
}

// Fonction pour créer le jeu en mode difficile
func (jeu *Engine) NouveauJeuDifficile() {
	// Charger les mots
	mots = ChargerMotsDepuisFichierHard()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	jeu.MotADeviner = motSansAccents
	jeu.LettresaTrouvées = make([]string, 0)
	for _, r := range jeu.MotADeviner {
		jeu.LettresaTrouvées = append(jeu.LettresaTrouvées, string(r))
	}
	jeu.LettresProposées = []string{}
	jeu.MotProposés = []string{}

	// Initialise les lettres à trouvées à "_"
	for i := 0; i < len(jeu.LettresaTrouvées); i++ {
		if jeu.LettresaTrouvées[i] == " " {
			jeu.LettresaTrouvées[i] = " "
		} else {
			jeu.LettresaTrouvées[i] = "_"
		}
	}
}

func (jeu *Engine) NouveauJeuBonus() {
	// Charger les mots
	//mots = ChargerMotsDepuisFichierBonus()

	// Choisir un mot aléatoire
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	motAleatoire := mots[rng.Intn(len(mots))]

	// Enlever les majuscules du mot
	motSansMajuscules := enleverMajuscules(motAleatoire)

	// Enlever les accents du mot
	motSansAccents := enleverAccents(motSansMajuscules)

	jeu.MotADeviner = motSansAccents
	jeu.LettresaTrouvées = make([]string, 0)
	for _, r := range jeu.MotADeviner {
		jeu.LettresaTrouvées = append(jeu.LettresaTrouvées, string(r))
	}
	jeu.LettresProposées = []string{}
	jeu.MotProposés = []string{}

	// Initialise les lettres à trouvées à "_"
	for i := 0; i < len(jeu.LettresaTrouvées); i++ {
		if jeu.LettresaTrouvées[i] == " " {
			jeu.LettresaTrouvées[i] = " "
		} else {
			jeu.LettresaTrouvées[i] = "_"
		}
	}
}
