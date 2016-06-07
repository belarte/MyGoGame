package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Four enemies
// Three enemies are within range
// Two are visible
// Get to the closest
func TestIntegration(t *testing.T) {
	char1 := &character.Fake{
		FakeVisibility:        character.DefaultVisibility,
		FakeRange:             character.DefaultRange,
		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 8, Y: 8}},
		FakeMovePointsComponent: character.FakeMovePointsComponent{
			FakeMovePoints: 10,
			FakeConsumeMP:  true,
		},
	}
	char2 := &character.Fake{
		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 3, Y: 9}},
	}
	char3 := &character.Fake{
		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 8, Y: 5}},
	}
	char4 := &character.Fake{
		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 13, Y: 13}},
	}
	char5 := &character.Fake{
		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 18, Y: 3}},
	}

	lvl := level.New(utils.Coord{X: 19, Y: 17}, 2)
	lvl.Map().SetCell(utils.Coord{X: 7, Y: 6}, level.WALL)
	lvl.Map().SetCell(utils.Coord{X: 8, Y: 6}, level.WALL)
	lvl.Map().SetCell(utils.Coord{X: 9, Y: 6}, level.WALL)
	lvl.AddCharacter(char1, utils.Coord{X: 8, Y: 8}, 0)
	lvl.AddCharacter(char2, utils.Coord{X: 3, Y: 9}, 1)
	lvl.AddCharacter(char3, utils.Coord{X: 8, Y: 5}, 1)
	lvl.AddCharacter(char4, utils.Coord{X: 13, Y: 13}, 1)
	lvl.AddCharacter(char5, utils.Coord{X: 18, Y: 3}, 1)

	c := newContext(lvl, char1)

	task := NewSequence(c)
	task.Add(NewGetVisibleEnemies(c))
	task.Add(NewGetClosestEnemies(c))
	task.Add(NewGetVantagePoint(c))
	//	task.Add(NewMoveToDestination(c)) TODO maybe fix this?

	if !task.CheckConditions() {
		t.Error("Tasks should have been checked ok")
	}

	if !task.Perform() {
		t.Error("Tasks should have been performed")
	}

	expectedVisible := 2
	expectedClosestEnemyPosition := utils.Coord{X: 3, Y: 9}
	expectedDestination := utils.Coord{X: 4, Y: 8}

	if len(c.visibleEnemies) != expectedVisible {
		t.Errorf("Wrong number of visible enemies\nresult=%d, expected=%d", len(c.visibleEnemies), expectedVisible)
	}

	if c.closestEnemyPosition != expectedClosestEnemyPosition {
		t.Errorf("Wrong enemy position\nresult=%+v, expected=%+v", c.closestEnemyPosition, expectedClosestEnemyPosition)
	}

	if c.destination != expectedDestination {
		t.Errorf("Wrong destination\nresult=%+v, expected=%+v", c.destination, expectedDestination)
	}
}
