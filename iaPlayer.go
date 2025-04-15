package main

import (
	"fmt"
)

type IaPlayer struct {
	game *Game
	model *Nnnarv
}

func (h *IaPlayer) Init(game *Game) {
	h.game = game
}

func (h IaPlayer) YourTurn() {

}

func (h IaPlayer) Click(y int) {
	fmt.Println("C'est à l'ia de jouer")
}
