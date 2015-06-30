package engine

type node struct {
	f_cost, g_cost, h_cost int
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

func (self *PathFinder) ShortestPath(start, dest Coord) path {
	self.openList = make(map[Coord]node)
	self.closedList = make(map[Coord]node)

	startNode := node{0, 0, 0, start}
	current := start

	self.openList[current] = startNode
	self.addToCloseList(current)
	self.addAdjacentCells(current, dest)

	for !EqualCoord(current, dest) && len(self.openList) > 0 {
		current = self.bestNode(self.openList)
		self.addToCloseList(current)
		self.addAdjacentCells(current, dest)
	}

	var result path

	if EqualCoord(current, dest) && !EqualCoord(start, dest) {
		result = self.retrievePath(start, dest)
	}

	return result
}

func (self *PathFinder) isInList(pos Coord, list nodeList) bool {
	for p, _ := range list {
		if EqualCoord(pos, p) {
			return true
		}
	}

	return false
}

func (self *PathFinder) addAdjacentCells(c, dest Coord) {
	coords := self.getAdjacentCells(c)

	for _, coord := range coords {
		if !self.isInList(coord, self.closedList) &&
			self.level.maps.GetCell(coord) != WALL &&
			!self.level.IsCharacterAtPosition(coord) {

			cellWeight := CellWeight(self.level.maps.GetCell(coord))
			dist := distance(coord, c)
			gcost := self.closedList[c].g_cost + dist*cellWeight
			hcost := distance(coord, dest)
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
	var result []Coord
	xx := make([]int, 0, 2)
	yy := make([]int, 0, 2)
	size := self.level.maps.Size()

	if c.x > 0 {
		xx = append(xx, c.x-1)
	}
	if c.x < size.x-1 {
		xx = append(xx, c.x+1)
	}
	if c.y > 0 {
		yy = append(yy, c.y-1)
	}
	if c.y < size.y-1 {
		yy = append(yy, c.y+1)
	}

	for _, x := range xx {
		result = append(result, Coord{x, c.y})
	}
	for _, y := range yy {
		result = append(result, Coord{c.x, y})
	}

	return result
}

func (self *PathFinder) bestNode(list nodeList) Coord {
	var result Coord
	cost := 0xFFFFFF

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

func (self *PathFinder) retrievePath(start, dest Coord) path {
	var result path

	tmp := self.closedList[dest]
	prev := tmp.parent
	weight := CellWeight(self.level.maps.GetCell(dest))
	result.add(dest, float64(weight))

	for !EqualCoord(prev, start) {
		weight = CellWeight(self.level.maps.GetCell(prev))
		result.add(prev, float64(weight))

		tmp = self.closedList[tmp.parent]
		prev = tmp.parent
	}

	result.reverse()

	return result
}
