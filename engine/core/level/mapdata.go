package level

import "math"

// Differents values for CellType.
var (
	NormalCell    = &NormalCellType{}
	DifficultCell = &DifficultCellType{}
	WallCell      = &WallCellType{}
)

// CellType represents the type of a cell in a Map.
type CellType interface {
	Cost() float64
}

// NormalCellType defines a normal cell type.
type NormalCellType struct{}

// Cost returns the cost of a normal cell type.
func (t NormalCellType) Cost() float64 {
	return 1.
}

// DifficultCellType defines a difficult cell type.
type DifficultCellType struct{}

// Cost returns the cost of a difficult cell type.
func (t DifficultCellType) Cost() float64 {
	return 2.
}

// WallCellType defines a wall cell type.
type WallCellType struct{}

// Cost returns the cost of wall cell type.
func (t WallCellType) Cost() float64 {
	return math.MaxFloat64
}
