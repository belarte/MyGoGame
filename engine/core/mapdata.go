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
		return 123456789.0
	}
}

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

func CompareEpsilon(left, right float64) bool {
	return left-right < 0.000001
}

/**
 * path class for pathFinder and MoveAction
 */
type pathStep struct {
	Coord Coord
	Cost  float64
}

type Path struct {
	Path []pathStep
}

func (self *Path) Add(coord Coord, cost float64) {
	self.Path = append(self.Path, pathStep{coord, cost})
}

func (self *Path) Size() int {
	return len(self.Path)
}

func (self *Path) Cost() float64 {
	result := 0.0
	for _, step := range self.Path {
		result += step.Cost
	}

	return result
}

func (self *Path) Reverse() {
	for i, j := 0, len(self.Path)-1; i < j; i, j = i+1, j-1 {
		self.Path[i], self.Path[j] = self.Path[j], self.Path[i]
	}
}
