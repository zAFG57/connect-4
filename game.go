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
	g.NextPlayerToPlay()
	g.UpdateGraphicalBoard()
}

func (g *Game) Click(y int) {
	if g.isToPlayer1ToPlay {
		g.player1.Click(y)
		return
	}
	g.player2.Click(y)
}

func (g *Game) NextPlayerToPlay() {
	g.isToPlayer1ToPlay = !g.isToPlayer1ToPlay
	if (g.isToPlayer1ToPlay) {
		go g.player1.YourTurn()
		return
	}
	go g.player2.YourTurn()
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
		wincube := getFoorConnected(g.board)
		drawWiningCube(wincube,g.img)
		if g.isToPlayer1ToPlay {
			fmt.Println("le joueur bleau a gagné")
			return
		}
		fmt.Println("le joueur rouge a gagné")
		return
	}
	g.NextPlayerToPlay()
}

func (g *Game) DrawPreviewCube(x int) {
	if !isGameRuning {
		return
	}
	for i:=0; i<7; i++ {
		if g.board[x][i] == 1 {
			g.UpdateGraphicalBoard()
			drawcube(g.img,x,i,color.RGBA{0,0,150,255})
			break
		}
	}
}

func drawWiningCube(winCube [4][2]uint8, img *image.RGBA) {
	for i:=0; i<4; i++ {
		for a:=0; a<40; a++ {
			for b:=0; b<40; b++ {
				if b>=5 && b<35 && a>=5 && a<35 {
					continue
				}
				img.Set(int(winCube[i][0])*40+a,280-(int(winCube[i][1])*40+b), color.RGBA{0,255,0,255})
			}
		}
	}
}

func drawcube(img *image.RGBA, x int, y int,color color.RGBA) {
	for a:=0; a<30; a++ {
		for b:=0; b<30; b++ {
			img.Set(x*40+a+5,280-(y*40+b+5), color)
		}
	}
}

func (g *Game) UpdateGraphicalBoard() {
	for x:=0; x<7; x++ {
		for y:=0; y<7; y++ {
			switch g.board[x][y] {
			case 0:
				drawcube(g.img,x,y,color.RGBA{0, 0, 255, 255})
			case 1:
				drawcube(g.img,x,y,color.RGBA{50, 50, 50, 255})
			case 2:
				drawcube(g.img,x,y,color.RGBA{255, 0, 0, 255})
			}
		}
	}
}

func (g *Game) IsFoorConnected() bool {
	return isFoorConnected(g.board)
}

func isFoorConnected(board [][]uint8) bool {
	for x:=0; x<7; x++ {
		for y:=0; y<4; y++ {
			if x<4 {
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
	return false
}

func getFoorConnected(board [][]uint8) [4][2]uint8 {
	for x:=0; x<7; x++ {
		for y:=0; y<4; y++ {
			if x<4 {
				if (
					board[x][y+3] == board[1+x][y+2] &&
					board[1+x][y+2] == board[2+x][y+1] &&
					board[2+x][y+1] == board[x+3][y] &&
					board[x+3][y] != 1) {
					return [4][2]uint8{
						{uint8(x), uint8(y+3)},
						{uint8(x+1), uint8(y+2)},
						{uint8(x+2), uint8(y+1)},
						{uint8(x+3), uint8(y)}}
				}
				if (
					board[x][y] == board[1+x][y+1] &&
					board[1+x][y+1] == board[2+x][y+2] &&
					board[2+x][y+2] == board[3+x][y+3] &&
					board[3+x][y+3] != 1) {
					return [4][2]uint8{
						{uint8(x),uint8(y)},
						{uint8(x+1),uint8(y+1)},
						{uint8(x+2),uint8(y+2)},
						{uint8(x+3),uint8(y+3)}}
				}
			}
			if (
				board[x][y] == board[x][1+y] &&
				board[x][1+y] == board[x][2+y] &&
				board[x][2+y] == board[x][3+y] &&
				board[x][3+y] != 1) {
					return [4][2]uint8{
						{uint8(x),uint8(y)},
						{uint8(x),uint8(y+1)},
						{uint8(x),uint8(y+2)},
						{uint8(x),uint8(y+3)}}
			}
			if (
				board[y][x] == board[1+y][x] &&
				board[1+y][x] == board[2+y][x] &&
				board[2+y][x] == board[3+y][x] &&
				board[3+y][x] != 1) {
					return [4][2]uint8{
						{uint8(y),uint8(x)},
						{uint8(y+1),uint8(x)},
						{uint8(y+2),uint8(x)},
						{uint8(y+3),uint8(x)}}
			}
		}
	}
	fmt.Println("error, il n'y a pas de gagnant")
	return [4][2]uint8{}
}

func isFoorWining(board [][]uint8, piece uint8) bool {
	for x:=0; x<7; x++ {
		for y:=0; y<4; y++ {
			if x<4 {
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

func getCopyOfBoard(board [][]uint8) [][]uint8 {
	ret := make([][]uint8, 7)
	for i := 0; i < 7; i++ {
		ret[i] = make([]uint8, 7)
		copy(ret[i], board[i])
	}
	return ret
}