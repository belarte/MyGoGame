package utils

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
