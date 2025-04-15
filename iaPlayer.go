package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type IaPlayer struct {
	game *Game
	model *Nnnarv
}

func (h *IaPlayer) Init(game *Game) {
	h.game = game
}

func (h *IaPlayer) YourTurn() {
	coord := make([]float64,7*7)
	for i:= 0; i<7; i++ {
		for y:= 0; y<7; y++ {
			coord[i*7+y] = float64(game.board[i][y])
		}
	}
	
	v,_ := strconv.ParseUint(h.model.GetValueOfPoint(coord,8),10,64)
	if (h.game.CheckIfIsValid(uint8(v))) {
		h.game.Play(uint8(v))
	} else {
		fmt.Println("l'IA ne peut pas jouer en: ",v)
		i := uint8(rand.Intn(8))
		for !h.game.CheckIfIsValid(i) {
			i = uint8(rand.Intn(8))
		}
		h.game.Play(i)
	}
}

func (h IaPlayer) Click(y int) {
	fmt.Println("C'est à l'ia de jouer")
}
