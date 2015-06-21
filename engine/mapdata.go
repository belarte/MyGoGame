package engine

const (
	NORMAL = iota
	DIFFICULT
	WALL
)

type CellType uint

func CellWeight(t CellType) int {
	switch t {
	case NORMAL:
		return 1
	case DIFFICULT:
		return 2
	default:
		return 0xFFFFFF
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

func distance(from, to Coord) int {
	// Compute the Manhattan distance
	x := from.x - to.x
	y := from.y - to.y
	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

/**
 * path class for pathFinder and MoveAction
 */
type pathStep struct {
	coord Coord
	cost  int
}

type path struct {
	path []pathStep
}

func (self *path) add(coord Coord, cost int) {
	self.path = append(self.path, pathStep{coord, cost})
}

func (self *path) size() int {
	return len(self.path)
}

func (self *path) cost() int {
	result := 0
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
