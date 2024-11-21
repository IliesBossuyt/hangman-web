package engine

// Structure du jeu
type Jeu struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
<<<<<<< Updated upstream:Hangman/struct.go
	EtapesPendu      []string
}
=======
	EtapesPendu      string
	Message          string
	Musique          bool
}
>>>>>>> Stashed changes:server/function/engine.go
