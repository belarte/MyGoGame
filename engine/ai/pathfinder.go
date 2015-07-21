package ai

import (
	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/utils"
	"math"
)

type node struct {
	fCost, gCost, hCost float64
	parent              utils.Coord
}

type nodeList map[utils.Coord]node

// PathFinder compute the shortest path betwin two coordinates.
// It handles the topography of the map.
type PathFinder struct {
	level                *core.Level
	closedList, openList nodeList
}

// NewPathFinder constructs a new PathFinder
func NewPathFinder(level *core.Level) *PathFinder {
	return &PathFinder{level, make(map[utils.Coord]node), make(map[utils.Coord]node)}
}

// ShortestPath computes the shortest path.
func (finder *PathFinder) ShortestPath(start, dest utils.Coord) core.Path {
	finder.openList = make(map[utils.Coord]node)
	finder.closedList = make(map[utils.Coord]node)

	startNode := node{0, 0, 0, start}
	current := start

	finder.openList[current] = startNode
	finder.addToCloseList(current)
	finder.addAdjacentCells(current, dest)

	for current != dest && len(finder.openList) > 0 {
		current = finder.bestNode(finder.openList)
		finder.addToCloseList(current)
		finder.addAdjacentCells(current, dest)
	}

	var result core.Path

	if current == dest && start != dest {
		result = finder.retrievePath(start, dest)
	}

	return result
}

func (finder *PathFinder) isInList(pos utils.Coord, list nodeList) bool {
	for p := range list {
		if pos == p {
			return true
		}
	}

	return false
}

func (finder *PathFinder) addAdjacentCells(c, dest utils.Coord) {
	coords := finder.getAdjacentCells(c)

	for _, coord := range coords {
		if !finder.isInList(coord, finder.closedList) &&
			finder.level.Map().GetCell(coord) != core.WALL &&
			!finder.level.IsCharacterAtPosition(coord) {

			cellWeight := core.CellWeight(finder.level.Map().GetCell(coord))
			dist := utils.Distance(coord, c)
			gcost := finder.closedList[c].gCost + dist*cellWeight
			hcost := utils.Distance(coord, dest)
			fcost := gcost + hcost
			tmp := node{fcost, gcost, hcost, c}

			if finder.isInList(coord, finder.openList) {
				if fcost < finder.openList[coord].fCost {
					finder.openList[coord] = tmp
				}
			} else {
				finder.openList[coord] = tmp
			}
		}
	}
}

func (finder *PathFinder) getAdjacentCells(c utils.Coord) []utils.Coord {
	size := finder.level.Map().Size()

	xx := make([]int, 0, 3)
	xx = append(xx, c.X)

	if c.X > 0 {
		xx = append(xx, c.X-1)
	}
	if c.X < size.X-1 {
		xx = append(xx, c.X+1)
	}

	yy := make([]int, 0, 3)
	yy = append(yy, c.Y)

	if c.Y > 0 {
		yy = append(yy, c.Y-1)
	}
	if c.Y < size.Y-1 {
		yy = append(yy, c.Y+1)
	}

	var result []utils.Coord
	for _, x := range xx {
		for _, y := range yy {
			if !(x == c.X && y == c.Y) {
				result = append(result, utils.Coord{x, y})
			}
		}
	}

	return result
}

func (finder *PathFinder) bestNode(list nodeList) utils.Coord {
	var result utils.Coord
	cost := math.MaxFloat64

	for c, n := range list {
		if n.fCost < cost {
			cost = n.fCost
			result = c
		}
	}

	return result
}

func (finder *PathFinder) addToCloseList(c utils.Coord) {
	n := finder.openList[c]
	finder.closedList[c] = n
	delete(finder.openList, c)
}

func (finder *PathFinder) retrievePath(start, dest utils.Coord) core.Path {
	var result core.Path

	tmp := finder.closedList[dest]
	current := dest
	previous := tmp.parent

	for current != start {
		weight := core.CellWeight(finder.level.Map().GetCell(current)) * utils.Distance(current, previous)
		result.Add(current, weight)

		current = previous
		tmp = finder.closedList[previous]
		previous = tmp.parent
	}

	result.Reverse()

	return result
}
