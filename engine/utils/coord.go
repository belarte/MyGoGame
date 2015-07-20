package utils

import (
	"math"
)

/**
 * class Coord
 * Base class for Map
 */
type Coord struct {
	X, Y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

var (
	NilCoord = Coord{-1, -1}
)

func Distance(from, to Coord) float64 {
	// Compute Euclidian distance
	x := float64(from.X - to.X)
	y := float64(from.Y - to.Y)

	return math.Sqrt(x*x + y*y)
}

// AreAdjacent checks if two Coord are adjacent.
func AreAdjacent(from, to Coord) bool {
	return Distance(from, to) < 1.5
}
