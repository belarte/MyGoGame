package core

import (
	"github.com/belarte/MyGoGame/engine/utils"
)

type pathStep struct {
	Coord utils.Coord
	Cost  float64
}

// Path is a list of adjacent weighted Coords.
type Path struct {
	Path []pathStep
}

// Add adds a weighted Coord to the path.
func (path *Path) Add(coord utils.Coord, cost float64) {
	if path.Size() == 0 || utils.AreAdjacent(path.Path[len(path.Path)-1].Coord, coord) {
		path.Path = append(path.Path, pathStep{coord, cost})
	}
}

// Size return the number of steps in the path.
func (path *Path) Size() int {
	return len(path.Path)
}

// Cost return the aggregated sum of the costs.
func (path *Path) Cost() float64 {
	result := 0.0
	for _, step := range path.Path {
		result += step.Cost
	}

	return result
}

// Reverse reverses the path.
func (path *Path) Reverse() {
	for i, j := 0, len(path.Path)-1; i < j; i, j = i+1, j-1 {
		path.Path[i], path.Path[j] = path.Path[j], path.Path[i]
	}
}
