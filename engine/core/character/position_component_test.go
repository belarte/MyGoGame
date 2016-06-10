package character

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/utils"
)

var expected = []utils.Coord{
	utils.Coord{X: 1, Y: 2},
	utils.Coord{X: 3, Y: 2},
	utils.Coord{X: 1, Y: 4},
	utils.Coord{X: 5, Y: 4},
	utils.Coord{X: 1, Y: 6},
	utils.Coord{X: 7, Y: 6},
}

func Test2DPosition(t *testing.T) {
	for _, pos := range expected {
		component := Position2DComponent{pos}
		if component.Position() != pos {
			t.Errorf("Wrong position, expected %v, got %v", pos, component.Position())
		}
	}
}

func Test2MoveTo(t *testing.T) {
	component := Position2DComponent{utils.NilCoord}

	for _, pos := range expected {
		component.MoveTo(pos)
		if component.Position() != pos {
			t.Errorf("Wrong position, expected %v, got %v", pos, component.Position())
		}
	}
}

func Test2DIsAtPosition(t *testing.T) {
	for _, pos := range expected {
		component := Position2DComponent{pos}
		if !component.IsAtPosition(pos) {
			t.Errorf("Wrong position, expected %v, got %v", pos, component.Position())
		}
	}
}
