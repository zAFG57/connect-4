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
	g.isToPlayer1ToPlay = true
	g.player1.YourTurn()
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
		go g.player1.YourTurn()
		return
	}
	go g.player2.YourTurn()
}


func (g *Game) UpdateGraphicalBoard() {
	for x:=0; x<7; x++ {
		for y:=0; y<7; y++ {
			switch g.board[x][y] {
			case 0:
				for a:=0; a<30; a++ {
					for b:=0; b<30; b++ {
						g.img.Set(x*40+a+5,280-(y*40+b+5), color.RGBA{0, 0, 255, 255})
					}
				}
			case 2:
				for a:=0; a<30; a++ {
					for b:=0; b<30; b++ {
						g.img.Set(x*40+a+5,280-(y*40+b+5), color.RGBA{255, 0, 0, 255})
					}
				}
			}
		}
	}
}

func (g *Game) IsFoorConnected() bool {
	return isFoorConnected(g.board)
}

func isFoorConnected(board [][]uint8) bool {
	for x:=0; x<7; x++ {
		for y:=0; y<3; y++ {
			if (
				board[x][y] == board[x][1+y] &&
				board[x][1+y] == board[x][2+y] &&
				board[x][2+y] == board[x][3+y] &&
				board[x][3+y] != 1) {
				return true
			}
			if (
				board[y][x] == board[1+y][x] &&
				board[1+y][x] == board[2+y][x] &&
				board[2+y][x] == board[3+y][x] &&
				board[3+y][x] != 1) {
				return true
			}
		}
	}
	for x:=0; x<4; x++ {
		for y:=0; y<4; y++ {
			if (
				board[x][y+3] == board[1+x][y+2] &&
				board[1+x][y+2] == board[2+x][y+1] &&
				board[2+x][y+1] == board[x+3][y] &&
				board[x+3][y] != 1) {
				return true
			}
			if (
				board[x][y] == board[1+x][y+1] &&
				board[1+x][y+1] == board[2+x][y+2] &&
				board[2+x][y+2] == board[3+x][y+3] &&
				board[3+x][y+3] != 1) {
				return true
			}
		}
	}
	return false
}

func isFoorWining(board [][]uint8, piece uint8) bool {
	for x:=0; x<7; x++ {
		for y:=0; y<3; y++ {
			if (
				board[x][y] == board[x][1+y] &&
				board[x][1+y] == board[x][2+y] &&
				board[x][2+y] == board[x][3+y] &&
				board[x][3+y] == piece) {
				return true
			}
			if (
				board[y][x] == board[1+y][x] &&
				board[1+y][x] == board[2+y][x] &&
				board[2+y][x] == board[3+y][x] &&
				board[3+y][x] == piece) {
				return true
			}
		}
	}
	for x:=0; x<4; x++ {
		for y:=0; y<4; y++ {
			if (
				board[x][y+3] == board[1+x][y+2] &&
				board[1+x][y+2] == board[2+x][y+1] &&
				board[2+x][y+1] == board[x+3][y] &&
				board[x+3][y] == piece) {
				return true
			}
			if (
				board[x][y] == board[1+x][y+1] &&
				board[1+x][y+1] == board[2+x][y+2] &&
				board[2+x][y+2] == board[3+x][y+3] &&
				board[3+x][y+3] == piece) {
				return true
			}
		}
	}
	return false
}

func (g *Game) CheckIfIsValid(n uint8) bool {
	return g.board[n][6] == 1
}

func getValidPlay(board [][]uint8) []uint8 {
	ret := make([]uint8,0)
	for i := uint8(0); i<7; i++ {
		if board[i][6] == 1 {
			ret = append(ret, i)
		}
	}
	return ret
}

func (g *Game) GetCopyOfBoard()[][]uint8 {
	return getCopyOfBoard(g.board)
}

func getCopyOfBoard(board [][]uint8) [][]uint8 {
	ret := make([][]uint8, 7)
	for i := 0; i < 7; i++ {
		ret[i] = make([]uint8, 7)
		copy(ret[i], board[i])
	}
	return ret
}