package main

import (
	e "engine/server/function"
)

func main() {
	var jeu e.Engine
	e.Run(&jeu)
}