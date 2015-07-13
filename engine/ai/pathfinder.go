package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
)

type node struct {
	f_cost, g_cost, h_cost float64
	parent                 Coord
}

type nodeList map[Coord]node

type PathFinder struct {
	level                *Level
	closedList, openList nodeList
}

func NewPathFinder(level *Level) *PathFinder {
	return &PathFinder{level, make(map[Coord]node), make(map[Coord]node)}
}

func (self *PathFinder) ShortestPath(start, dest Coord) Path {
	self.openList = make(map[Coord]node)
	self.closedList = make(map[Coord]node)

	startNode := node{0, 0, 0, start}
	current := start

	self.openList[current] = startNode
	self.addToCloseList(current)
	self.addAdjacentCells(current, dest)

	for current != dest && len(self.openList) > 0 {
		current = self.bestNode(self.openList)
		self.addToCloseList(current)
		self.addAdjacentCells(current, dest)
	}

	var result Path

	if current == dest && start != dest {
		result = self.retrievePath(start, dest)
	}

	return result
}

func (self *PathFinder) isInList(pos Coord, list nodeList) bool {
	for p, _ := range list {
		if pos == p {
			return true
		}
	}

	return false
}

func (self *PathFinder) addAdjacentCells(c, dest Coord) {
	coords := self.getAdjacentCells(c)

	for _, coord := range coords {
		if !self.isInList(coord, self.closedList) &&
			self.level.Map().GetCell(coord) != WALL &&
			!self.level.IsCharacterAtPosition(coord) {

			cellWeight := CellWeight(self.level.Map().GetCell(coord))
			dist := Distance(coord, c)
			gcost := self.closedList[c].g_cost + dist*cellWeight
			hcost := Distance(coord, dest)
			fcost := gcost + hcost
			tmp := node{fcost, gcost, hcost, c}

			if self.isInList(coord, self.openList) {
				if fcost < self.openList[coord].f_cost {
					self.openList[coord] = tmp
				}
			} else {
				self.openList[coord] = tmp
			}
		}
	}
}

func (self *PathFinder) getAdjacentCells(c Coord) []Coord {
	size := self.level.Map().Size()

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

	var result []Coord
	for _, x := range xx {
		for _, y := range yy {
			if !(x == c.X && y == c.Y) {
				result = append(result, Coord{x, y})
			}
		}
	}

	return result
}

func (self *PathFinder) bestNode(list nodeList) Coord {
	var result Coord
	cost := 123456789.0

	for c, n := range list {
		if n.f_cost < cost {
			cost = n.f_cost
			result = c
		}
	}

	return result
}

func (self *PathFinder) addToCloseList(c Coord) {
	n := self.openList[c]
	self.closedList[c] = n
	delete(self.openList, c)
}

func (self *PathFinder) retrievePath(start, dest Coord) Path {
	var result Path

	tmp := self.closedList[dest]
	current := dest
	previous := tmp.parent

	for current != start {
		weight := CellWeight(self.level.Map().GetCell(current)) * Distance(current, previous)
		result.Add(current, weight)

		current = previous
		tmp = self.closedList[previous]
		previous = tmp.parent
	}

	result.Reverse()

	return result
}
