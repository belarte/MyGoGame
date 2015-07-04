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

func EqualCoord(left, right Coord) bool {
	return left.X == right.X && left.Y == right.Y
}

func Distance(from, to Coord) float64 {
	// Compute Euclidian distance
	x := float64(from.X - to.X)
	y := float64(from.Y - to.Y)

	return math.Sqrt(x*x + y*y)
}
