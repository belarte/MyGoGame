package engine

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
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

func EqualCoord(left, right Coord) bool {
	return left.x == right.x && left.y == right.y
}

func Distance(from, to Coord) float64 {
	// Compute Euclidian distance
	x := float64(from.x - to.x)
	y := float64(from.y - to.y)

	return math.Sqrt(x*x + y*y)
}

func CompareEpsilon(left, right float64) bool {
	return left-right < 0.000001
}

/**
 * path class for pathFinder and MoveAction
 */
type pathStep struct {
	coord Coord
	cost  float64
}

type path struct {
	path []pathStep
}

func (self *path) add(coord Coord, cost float64) {
	self.path = append(self.path, pathStep{coord, cost})
}

func (self *path) size() int {
	return len(self.path)
}

func (self *path) cost() float64 {
	result := 0.0
	for _, step := range self.path {
		result += step.cost
	}

	return result
}

func (self *path) reverse() {
	for i, j := 0, len(self.path)-1; i < j; i, j = i+1, j-1 {
		self.path[i], self.path[j] = self.path[j], self.path[i]
	}
}
