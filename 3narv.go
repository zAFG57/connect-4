package main

import (
	"fmt"
	"math"
)

type Nnnarv struct {
	subSpaceGestionaire SubSpaceGestionaireEvolv
	nbSubDiv 			int
	nbDimentions 		int
	minCoord 			float64
	maxCoord 			float64
	step 				float64
}

func (n *Nnnarv) Init(nbSubDiv int, nbDimentions int, minCoord float64, maxCoord float64) {
	fmt.Println("")
	n.nbSubDiv = nbSubDiv
	n.nbDimentions = nbDimentions
	n.minCoord = minCoord
	n.maxCoord = maxCoord
	n.step = (maxCoord - minCoord) / float64(nbSubDiv)
}

func (n *Nnnarv) GetSubSpaceFromCoord(coord []float64, needToExist bool) *SubSpace {
	return n.subSpaceGestionaire.GetSubSpaceFromCoord(coord, n, needToExist)
}

func (n *Nnnarv) AddPoint(p Point) {
	sb := n.GetSubSpaceFromCoord(p.coord, true)
	sb.AddPoint(p)
}

func (n *Nnnarv) GetSubSapceAround(coord []float64, distance int) []SubSpace {
	subSpace := make([]SubSpace, 0)
	sbCoord := FindCoordAround2(n.nbDimentions, distance)
	for i := 0; i < len(*sbCoord); i++ {
		tCoord, error := n.ApplySubSpaceCoord(coord, (*sbCoord)[i])
		if error {
			continue
		}
		subSpace = append(subSpace, *n.GetSubSpaceFromCoord(tCoord,false))
	}
	return subSpace
}

func (n *Nnnarv) ApplySubSpaceCoord(center []float64, coord []int) ([]float64,bool) {
	res := make([]float64, len(center))
	for i := 0; i < len(center); i++ {
		res[i] = center[i] + float64(coord[i]) * n.step
		if res[i] < n.minCoord || res[i] >= n.maxCoord {
			return nil,true
		}
	}
	return res,false
}

func (n *Nnnarv) getNNearestPoint(coord []float64, nbNearest int) ([]Point, []float64) {
	selected := make([]Point, 0)
	for nbAround := 1; len(selected) < nbNearest; nbAround++ {
		selected = make([]Point, 0)
		subSpaceFound := n.GetSubSapceAround(coord, nbAround)
		for i := 0; i < len(subSpaceFound); i++ {
			for j := 0; j < len(subSpaceFound[i].listPoint); j++ {
				if subSpaceFound[i].listPoint[j].GetDistanceFromCoord(coord) <= float64(nbAround)*n.step {
					selected = append(selected, subSpaceFound[i].listPoint[j])
				}
			}
		}
	}
	found := make([]Point, nbNearest)
	dist := make([]float64, nbNearest)
	idx := 0
	for i:=0; i < nbNearest; i++ {
		found[i] = selected[i]
		dist[i] = found[i].GetDistanceFromCoord(coord)
	}
	maxDist := math.MaxFloat64
	for j := 0; j < nbNearest; j++ {
		tDist := found[j].GetDistanceFromCoord(coord)
		if tDist > maxDist {
			maxDist = tDist
			idx = j
		}
	}
	for i := nbNearest; i < len(selected); i++ {
		tDist := selected[i].GetDistanceFromCoord(coord)
		if tDist < maxDist {
			found[idx] = selected[i]
			dist[idx] = tDist
			for j := 0; j < nbNearest; j++ {
				tDist := found[j].GetDistanceFromCoord(coord)
				if tDist > maxDist {
					maxDist = tDist
					idx = j
				}
			}
		}
	}
	return selected, dist
}
/*
func (n *Nnnarv) GetValueOfPoint(coord []float64, nbNearest int) []float64 {
	points, dists := n.getNNearestPoint(coord, nbNearest)
	ttPart := float64(0)
	res := make([]float64, len(points[0].value))
	for i := 0; i < len(points); i++ {
		if (dists[i] == 0) {
			copy(res,points[i].value)
			return res
		}
		ttPart += float64(1)/dists[i]
		for y:= 0; y<len(points[i].value); y++ {
			res[y] += points[i].value[y] / dists[i]
		}
	}
	for i := 0; i < len(res); i++ {
		res[i] /= ttPart
	}
	return res
}
*/
func (n *Nnnarv) GetValueOfPoint(coord []float64, nbNearest int) string {
	points, dists := n.getNNearestPoint(coord, nbNearest)

	label := make(map[string]float64)
	for i := 0; i < len(points); i++ {
		tvalue := points[i].valueStr
		if val, exists := label[tvalue]; exists {
			label[tvalue] = val + 1/(dists[i]+1)
		} else {
			label[tvalue] = 1/(dists[i]+1)
		}
	}
	maxVal := float64(0)
	var res string
	for k, v := range label {
		if v > maxVal {
			maxVal = v
			res = k
		}
	}
	return res
}