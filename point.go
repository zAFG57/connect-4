package main

import (
	"math"
)

type Point struct {
	coord 		[]float64
	// valueStr 	string
	value		[]float64
}

func (p1 *Point) GetDistance(p2 *Point) float64 {
	return p1.GetDistanceFromCoord(p2.coord)
}

func (p *Point) GetDistanceFromCoord(coord []float64) float64 {
	var sum float64
	for i := 0; i < len(p.coord); i++ {
		sum += (p.coord[i] - coord[i]) * (p.coord[i] - coord[i])
	}
	sum = math.Sqrt(sum)
	return sum
}