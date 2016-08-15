package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestGetClosestEnemyCheckConditionsNoVisibleEnemies(t *testing.T) {
	context := &context{}
	task := NewGetClosestEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetClosestEnemyCheckConditionsVisibleEnemies(t *testing.T) {
	context := &context{visibleEnemies: make([]charDist, 1, 1)}
	task := NewGetClosestEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetClosestEnemyPerformOneEnemy(t *testing.T) {
	position := utils.Coord{X: 0, Y: 1}
	char := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: position},
	}
	opponents := []charDist{charDist{char, 1}}
	context := &context{visibleEnemies: opponents}
	task := NewGetClosestEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}

	if context.closestEnemy != char {
		t.Errorf("Expected %+v, got %+v", char, context.closestEnemy)
	}

	if context.closestEnemyPosition != position {
		t.Errorf("Expected %+v, got %+v", position, context.closestEnemyPosition)
	}
}

func TestGetClosestEnemyPerformThreeEnemies(t *testing.T) {
	position1 := utils.Coord{X: 0, Y: 1}
	position2 := utils.Coord{X: 0, Y: 2}
	position3 := utils.Coord{X: 0, Y: 3}
	char1 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: position1},
	}
	char2 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: position2},
	}
	char3 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: position3},
	}

	opponents := []charDist{
		charDist{char3, 3},
		charDist{char1, 1},
		charDist{char2, 2},
	}
	context := &context{visibleEnemies: opponents}
	task := NewGetClosestEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}

	if context.closestEnemy != char1 {
		t.Errorf("Expected %v, got %+v", char1, context.closestEnemy)
	}

	if context.closestEnemyPosition != position1 {
		t.Errorf("Expected %+v, got %+v", position1, context.closestEnemyPosition)
	}
}
