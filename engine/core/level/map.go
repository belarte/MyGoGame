package level

import (
	"errors"
	"strconv"

	"github.com/belarte/MyGoGame/engine/utils"
)

// Map represents a grid of cells.
type Map struct {
	size utils.Coord
	grid []CellType
}

//NewMap returns the new Map.
func NewMap(size utils.Coord) *Map {
	grid := make([]CellType, size.X*size.Y)
	for i := range grid {
		grid[i] = &NormalCell
	}
	return &Map{size, grid}
}

// Size returns the size of the grid.
func (m *Map) Size() utils.Coord {
	return m.size
}

// GetCell returns the type of the given cell.
func (m *Map) GetCell(coord utils.Coord) CellType {
	index, _ := m.getIndex(coord)
	return m.grid[index]
}

// SetCell set the type of the given cell.
func (m *Map) SetCell(coord utils.Coord, t CellType) {
	index, _ := m.getIndex(coord)
	m.grid[index] = t
}

func (m *Map) getIndex(coord utils.Coord) (int, error) {
	if coord.X > m.size.X || coord.Y > m.size.Y {
		return 0, errors.New("Coordinates out of range: (" +
			strconv.Itoa(int(coord.X)) + ", " +
			strconv.Itoa(int(coord.Y)) + ")")
	}

	return coord.Y*m.size.X + coord.X, nil
}

// IsWithinBounds checks if the given cell is within the grid.
func (m *Map) IsWithinBounds(pos utils.Coord) bool {
	size := m.Size()
	return pos.X >= 0 && pos.X < size.X && pos.Y >= 0 && pos.Y < size.Y
}
