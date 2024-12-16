package engine

// Je crée ma structure
type Engine struct {
	MotADeviner      string
	LettresaTrouvées []string
	ViesRestantes    int
	LettresProposées []string
	MotProposés      []string
	EtapesPendu      string
	EtapesBonus      string
	Message          string
	Musique          bool
	Score            int
	Value            int
	MeilleurScore    int
}
