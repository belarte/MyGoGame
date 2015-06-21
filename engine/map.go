package engine

import (
	"errors"
	"strconv"
)

type Map struct {
	size Coord
	grid []CellType
}

func NewMap(size Coord) *Map {
	grid := make([]CellType, size.x*size.y)
	return &Map{size, grid}
}

func (self *Map) Size() Coord {
	return self.size
}

func (self *Map) GetCell(coord Coord) CellType {
	index, _ := self.getIndex(coord)
	return self.grid[index]
}

func (self *Map) SetCell(coord Coord, t CellType) {
	index, _ := self.getIndex(coord)
	self.grid[index] = t
}

func (self *Map) getIndex(coord Coord) (int, error) {
	if coord.x > self.size.x || coord.y > self.size.y {
		return 0, errors.New("Coordinates out of range: (" +
			strconv.Itoa(int(coord.x)) + ", " +
			strconv.Itoa(int(coord.y)) + ")")
	}

	return coord.y*self.size.x + coord.x, nil
}

func (self *Map) isWithinBounds(pos Coord) bool {
	size := self.Size()
	return size.x < pos.x && size.y < pos.y
}
