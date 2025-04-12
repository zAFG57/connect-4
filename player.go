package main


type Player interface {
	Init(game *Game)
	YourTurn()
	Click(y int)
}

type HumanPlayer struct {
	game *Game
}

func (h *HumanPlayer) Init(game *Game) {
	h.game = game
}

func (h HumanPlayer) YourTurn() {

}

func (h HumanPlayer) play(y int) {
	h.game.Play(uint8(y / 40))
}

func (h HumanPlayer) Click(y int) {
	h.play(y)
}
