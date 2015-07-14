package core

import (
	"errors"
	. "github.com/belarte/MyGoGame/engine/utils"
	"strconv"
)

type Map struct {
	size Coord
	grid []CellType
}

func NewMap(size Coord) *Map {
	grid := make([]CellType, size.X*size.Y)
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
	if coord.X > self.size.X || coord.Y > self.size.Y {
		return 0, errors.New("Coordinates out of range: (" +
			strconv.Itoa(int(coord.X)) + ", " +
			strconv.Itoa(int(coord.Y)) + ")")
	}

	return coord.Y*self.size.X + coord.X, nil
}

func (self *Map) IsWithinBounds(pos Coord) bool {
	size := self.Size()
	return pos.X >= 0 && pos.X < size.X && pos.Y >= 0 && pos.Y < size.Y
}
