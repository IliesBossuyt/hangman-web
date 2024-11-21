package engine

// Je crée ma structure
type Engine struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
	EtapesPendu      string
	Message          string
	Musique          bool
}