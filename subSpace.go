package main

import (
	
)

type SubSpace struct {
	listPoint []Point
}

func (s *SubSpace) AddPoint(p Point) {
	s.listPoint = append(s.listPoint, p)
}