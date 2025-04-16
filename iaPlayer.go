package main

import (
	"fmt"
	"math/rand"
	"math"
)

type IaPlayer struct {
	game *Game
}

func (ia *IaPlayer) Init(game *Game) {
	ia.game = game
}

func (ia IaPlayer) Click(y int) {
	fmt.Println("C'est à l'ia de jouer")
}

func (ia *IaPlayer) YourTurn() {
	v,_ := ia.MinMax(ia.game.board,8,math.MinInt,math.MaxInt,false)
	if (ia.game.CheckIfIsValid(uint8(v))) {
		ia.game.Play(uint8(v))
	} else {
		fmt.Println("l'IA ne peut pas jouer en: ",v)
		i := uint8(rand.Intn(7))
		for !ia.game.CheckIfIsValid(i) {
			i = uint8(rand.Intn(7))
		}
		ia.game.Play(i)
	}
}

func (ia *IaPlayer) MinMax(board [][]uint8,idx uint8, min int, max int, needMax bool) (uint8,int) {
	if isFoorWining(board,0) {
		return 8, math.MinInt
	}
	board = getCopyOfBoard(board)
	valid := getValidPlay(board)
	if isFoorConnected(board) || len(valid) == 0 {
		if isFoorWining(board,2) {
			return 8, math.MaxInt
		}
		return 8,0
	}
	if idx == 0 {
		return 8,ia.ScorePos(board,2,0)
	}
	if needMax {
		val := math.MinInt
		col := valid[0]
		r := 0
		for i:=0; i<len(valid); i++ {
			for y:=0; y<7; y++ {
				if (board[i][y] == 1) {
					board[i][y] = 2
					r = y
					break
				}
			}
			_,nv := ia.MinMax(board,idx-1,min,max,false)
			if nv>val {
				val = nv
				col = valid[i]
				if val>min {
					min = val
				}
				if min>=max {
					break
				}
			}
			board[i][r] = 1
		}
		return uint8(col),val
	}
	val := math.MaxInt
	col := valid[0]
	r:= 0
	for i:=0; i<len(valid); i++ {
		for y:=0; y<7; y++ {
			if (board[i][y] == 1) {
				board[i][y] = 0
				r = y
				break
			}
		}
		_,nv := ia.MinMax(board,idx-1,min,max,true)
		if nv<val {
			val = nv
			col = valid[i]
			if val<max {
				max = val
			}
			if min>=max {
				break
			}
		}
		board[i][r] = 1
	}
	return uint8(col),val
}

func eval(arr [4]uint8, piece uint8, rivPiece uint8) int {
	p :=0
	rivP :=0
	for i:=0; i<4; i++ {
		if arr[i] == piece {
			p++
			continue
		}
		if arr[i] == rivPiece {
			rivP++
		}
	}
	if p == 4 {
		return 1000
	}
	if rivP == 4 {
		return -1000
	}
	if p == 3 && rivP == 0 {
		return 5
	}
	if rivP == 3 && p == 0 {
		return -5
	}
	if p == 2 && rivP == 0{
		return 2
	}
	if rivP == 2 && p == 0{
		return -2
	}
	return 0
}

func (ia *IaPlayer) ScorePos(board [][]uint8, piece uint8,rivPiece uint8) int {
	ret := 0
	for x:=0; x<7; x++ {
		for y:=0; y<3; y++ {
			ret += eval(
				[4]uint8{board[x][y],board[x][1+y],board[x][2+y],board[x][3+y]},
			piece,rivPiece)
			ret += eval(
				[4]uint8{board[y][x],board[1+y][x],board[2+y][x],board[3+y][x]},
			piece,rivPiece)
		}
	}
	for x:=0; x<4; x++ {
		for y:=0; y<4; y++ {
			ret += eval(
				[4]uint8{board[x][y+3],board[1+x][y+2],board[2+x][y+1],board[x+3][y]},
			piece,rivPiece)
			ret += eval(
				[4]uint8{board[x][y],board[1+x][y+1],board[2+x][y+2],board[3+x][y+3]},
			piece,rivPiece)
		}
	}
	return ret
}