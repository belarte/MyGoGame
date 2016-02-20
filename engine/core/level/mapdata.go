package level

import (
	"math"
)

// Differents values for CellType.
const (
	NORMAL = iota
	DIFFICULT
	WALL
)

// CellType represents the type of a cell in a Map.
type CellType uint

// CellWeight returns the weight corresponding CellType
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
