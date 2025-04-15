package main

import (
	"math"
)

const MAX_DIMENSION = 10	// valable tant que nbSubDiv^10 < 2^31-1

type SubSpaceGestionaireEvolv struct {
	mapSubSpaceGestionaire		map[int]*SubSpaceGestionaireEvolv
	mapSubSpace 				map[int]*SubSpace
}

func (s *SubSpaceGestionaireEvolv) GetSubSpaceFromCoord(coord []float64, nnnarv *Nnnarv, needToExist bool) *SubSpace {
	isFinal := len(coord) <= MAX_DIMENSION
	idx := 0
	for i := 0; !(i >= MAX_DIMENSION) && !(i >= len(coord)); i++ {
		idx += int((coord[i] - nnnarv.minCoord) / nnnarv.step) * int(math.Pow(float64(nnnarv.nbSubDiv), float64(i)))
	}

	if isFinal {
		subspace, ok := s.mapSubSpace[idx]
		if !ok {
			if !needToExist {
				return &SubSpace{}
			}
			subspace = &SubSpace{}
			if len(s.mapSubSpace) == 0 {
				s.mapSubSpace = make(map[int]*SubSpace)
			}
			s.mapSubSpace[idx] = subspace
			return subspace
		}
		return subspace
	}

	subSpaceGestionaire, ok := s.mapSubSpaceGestionaire[idx]
	if !ok {
		if !needToExist {
			return &SubSpace{}
		}
		if len(s.mapSubSpaceGestionaire) == 0 {
			s.mapSubSpaceGestionaire = make(map[int]*SubSpaceGestionaireEvolv)
		}
		s.mapSubSpaceGestionaire[idx] = &SubSpaceGestionaireEvolv{}
		return s.mapSubSpaceGestionaire[idx].GetSubSpaceFromCoord(coord[MAX_DIMENSION:], nnnarv, needToExist)
	}
	return subSpaceGestionaire.GetSubSpaceFromCoord(coord[MAX_DIMENSION:], nnnarv, needToExist)
}