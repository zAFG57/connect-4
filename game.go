package main

import (
	"fmt"
	"image"
	"image/color"
)

type Game struct {
	player1				Player
	player2				Player
	isToPlayer1ToPlay	bool
	board				[][]uint8
	img					*image.RGBA
}

func (g *Game) Init(p1 Player, p2 Player, img *image.RGBA) {
	g.player1 = p1
	g.player2 = p2
	g.board = make([][]uint8, 7)
	g.img = img
	for i := 0; i < 7; i++ {
		g.board[i] = make([]uint8, 7)
		for y := 0; y < 7; y++ {
			g.board[i][y] = 1
		}
	}
}

func (g *Game) Click(y int) {
	if g.isToPlayer1ToPlay {
		g.player1.Click(y)
		return
	}
	g.player2.Click(y)
}

func (g *Game) Play(n uint8) {
	if !g.CheckIfIsValid(n) {
		fmt.Println("erreur, ce coup n'est pas valid")
		isGameRuning = false
		return
	}
	var x uint8 = 2
	if g.isToPlayer1ToPlay {
		x = 0
	}
	i:=0
	for ; g.board[n][i] != 1; i++ {}
	g.board[n][i] = x
	g.UpdateGraphicalBoard()
	if g.IsFoorConnected() {
		isGameRuning = false
		if g.isToPlayer1ToPlay {
			fmt.Println("le joueur bleau a gagné")
			return
		}
		fmt.Println("le joueur rouge a gagné")
		return
	}
	g.isToPlayer1ToPlay = !g.isToPlayer1ToPlay
	if (g.isToPlayer1ToPlay) {
		g.player1.YourTurn()
		return
	}
	g.player2.YourTurn()
}

func (g *Game) CheckIfIsValid(n uint8) bool {
	return g.board[n][6] == 1
}

func (g *Game) UpdateGraphicalBoard() {
	for x:=0; x<7; x++ {
		for y:=0; y<7; y++ {
			switch g.board[x][y] {
			case 0:
				for a:=0; a<40; a++ {
					for b:=0; b<40; b++ {
						g.img.Set(x*40+a,280-(y*40+b), color.RGBA{0, 0, 255, 255})
					}
				}
			case 2:
				for a:=0; a<40; a++ {
					for b:=0; b<40; b++ {
						g.img.Set(x*40+a,280-(y*40+b), color.RGBA{255, 0, 0, 255})
					}
				}
			}
		}
	}
}

func (g *Game) IsFoorConnected() bool {
	for x:=0; x<7; x++ {
		for y:=0; y<3; y++ {
			if (
				g.board[x][y] == g.board[x][1+y] &&
				g.board[x][1+y] == g.board[x][2+y] &&
				g.board[x][2+y] == g.board[x][3+y] &&
				g.board[x][3+y] != 1) {
				return true
			}
			if (
				g.board[y][x] == g.board[1+y][x] &&
				g.board[1+y][x] == g.board[2+y][x] &&
				g.board[2+y][x] == g.board[3+y][x] &&
				g.board[3+y][x] != 1) {
				return true
			}
		}
	}

	for x:=0; x<4; x++ {
		for y:=0; y<4; y++ {
			if (
				g.board[x][y+3] == g.board[1+x][y+2] &&
				g.board[1+x][y+2] == g.board[2+x][y+1] &&
				g.board[2+x][y+1] == g.board[x+3][y] &&
				g.board[x+3][y] != 1) {
				return true
			}

			if (
				g.board[x][y] == g.board[1+x][y+1] &&
				g.board[1+x][y+1] == g.board[2+x][y+2] &&
				g.board[2+x][y+2] == g.board[3+x][y+3] &&
				g.board[3+x][y+3] != 1) {
				return true
			}
		}
	}

	return false
}