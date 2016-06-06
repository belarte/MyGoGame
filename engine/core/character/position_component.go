package character

import "github.com/belarte/MyGoGame/engine/utils"

// PositionComponent interface
type PositionComponent interface {
	Position() utils.Coord
	MoveTo(utils.Coord)
	IsAtPosition(utils.Coord) bool
}

// Position2DComponent handle position in 2D space.
type Position2DComponent struct {
	position utils.Coord
}

// Position returns the current position.
func (component *Position2DComponent) Position() utils.Coord {
	return component.position
}

// MoveTo sets a new position.
func (component *Position2DComponent) MoveTo(position utils.Coord) {
	component.position = position
}

// IsAtPosition return true is the parameter match the current position.
func (component *Position2DComponent) IsAtPosition(position utils.Coord) bool {
	return component.position == position
}
