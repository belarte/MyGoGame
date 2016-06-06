package character

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/utils"
)

var expected = []utils.Coord{
	utils.Coord{1, 2},
	utils.Coord{3, 2},
	utils.Coord{1, 4},
	utils.Coord{5, 4},
	utils.Coord{1, 6},
	utils.Coord{7, 6},
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
