package core

import (
	"math"
)

const (
	NORMAL = iota
	DIFFICULT
	WALL
)

type CellType uint

func CellWeight(t CellType) float64 {
	switch t {
	case NORMAL:
		return 1.
	case DIFFICULT:
		return 2.
	default:
		return math.MaxFloat64
	}
}
