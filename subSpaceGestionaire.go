package main

import (
	"math"
)

type SubSpaceGestionaire struct {
	listSubSpaceGestionaire 	[] SubSpaceGestionaire
	listSubSpace 				[] SubSpace
	isFinal 					bool
	isInited 					bool
}

func (s *SubSpaceGestionaire) Init(nbDimentions int, nbEtage int, nbSubSpace int) {
	s.isInited = true
	s.isFinal = nbDimentions - nbEtage * 10 < 10
	if  s.isFinal {
		s.listSubSpace = make([]SubSpace,int(math.Pow(float64(nbSubSpace), float64(nbDimentions))))
		return
	}
	s.listSubSpaceGestionaire = make([]SubSpaceGestionaire, int(math.Pow(float64(nbSubSpace), float64(10))))
}

func (s *SubSpaceGestionaire) GetSubSpaceFromCoord(coord []float64, nnnarv *Nnnarv, needToExist bool) *SubSpace {
	if (!s.isInited && !needToExist) {
		return &SubSpace{}
	}
	if (needToExist && !s.isInited) {
		s.Init(nnnarv.nbDimentions,(nnnarv.nbDimentions-len(coord))/10,nnnarv.nbSubDiv)
	}
	if s.isFinal {
		idx := 0
		for i := 0; i < len(coord); i++ {
			idx += int((coord[i] - nnnarv.minCoord) / nnnarv.step) * int(math.Pow(float64(nnnarv.nbSubDiv), float64(i)))
		}
		return &s.listSubSpace[idx]
	}
	idx := 0
	for i := 0; i < 10; i++ {
		idx += int((coord[i] - nnnarv.minCoord) / nnnarv.step) * int(math.Pow(float64(nnnarv.nbSubDiv), float64(i)))
	}
	return s.listSubSpaceGestionaire[idx].GetSubSpaceFromCoord(coord[10:], nnnarv, needToExist)
}