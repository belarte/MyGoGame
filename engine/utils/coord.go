package utils

import (
	"math"
)

// Coord struct. It has visible fields for access purpose.
type Coord struct {
	X, Y int
}

var (
	// NilCoord is to be used as an undefined Coord.
	NilCoord = Coord{-1, -1}
)

// Distance compute the Euclidian distance between two Coords.
func Distance(from, to Coord) float64 {
	x := float64(from.X - to.X)
	y := float64(from.Y - to.Y)

	return math.Sqrt(x*x + y*y)
}

// AreAdjacent checks if two Coord are adjacent.
func AreAdjacent(from, to Coord) bool {
	return from != to && Distance(from, to) < 1.5
}
