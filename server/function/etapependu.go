package main

func (jeu *Engine) EtapePendu() {
	if jeu.ViesRestantes == 11 {
		jeu.EtapesPendu = "/static/1.png"
	}
	if jeu.ViesRestantes == 10 {
		jeu.EtapesPendu = "/static/2.png"
	}
	if jeu.ViesRestantes == 9 {
		jeu.EtapesPendu = "/static/3.png"
	}
	if jeu.ViesRestantes == 8 {
		jeu.EtapesPendu = "/static/4.png"
	}
	if jeu.ViesRestantes == 7 {
		jeu.EtapesPendu = "/static/5.png"
	}
	if jeu.ViesRestantes == 6 {
		jeu.EtapesPendu = "/static/6.png"
	}
	if jeu.ViesRestantes == 5 {
		jeu.EtapesPendu = "/static/7.png"
	}
	if jeu.ViesRestantes == 4 {
		jeu.EtapesPendu = "/static/8.png"
	}
	if jeu.ViesRestantes == 3 {
		jeu.EtapesPendu = "/static/9.png"
	}
	if jeu.ViesRestantes == 2 {
		jeu.EtapesPendu = "/static/10.png"
	}
	if jeu.ViesRestantes == 1 {
		jeu.EtapesPendu = "/static/11.png"
	}
}
