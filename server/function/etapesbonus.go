package engine

// Je définit les étapes du pendu
func (jeu *Engine) EtapeBonus() {
	if jeu.ViesRestantes == 7 {
		jeu.EtapesBonus = "/static/img/etapesbonus/1.png"
	}
	if jeu.ViesRestantes == 6 {
		jeu.EtapesBonus = "/static/img/etapesbonus/2.png"
	}
	if jeu.ViesRestantes == 5 {
		jeu.EtapesBonus = "/static/img/etapesbonus/3.png"
	}
	if jeu.ViesRestantes == 4 {
		jeu.EtapesBonus = "/static/img/etapesbonus/4.png"
	}
	if jeu.ViesRestantes == 3 {
		jeu.EtapesBonus = "/static/img/etapesbonus/5.png"
	}
	if jeu.ViesRestantes == 2 {
		jeu.EtapesBonus = "/static/img/etapesbonus/6.png"
	}
	if jeu.ViesRestantes == 1 {
		jeu.EtapesBonus = "/static/img/etapesbonus/7.png"
	}
}
