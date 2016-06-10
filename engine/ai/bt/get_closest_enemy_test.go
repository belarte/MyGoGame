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
	context := &context{visibleEnemies: make([]charPosDist, 1, 1)}
	task := NewGetClosestEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetClosestEnemyPerformOneEnemy(t *testing.T) {
	char := &character.Fake{}
	position := utils.Coord{X: 0, Y: 1}
	opponents := []charPosDist{charPosDist{char, position, 1}}
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
	char1 := &character.Fake{}
	char2 := &character.Fake{}
	char3 := &character.Fake{}
	position1 := utils.Coord{X: 0, Y: 1}
	position2 := utils.Coord{X: 0, Y: 2}
	position3 := utils.Coord{X: 0, Y: 3}

	opponents := []charPosDist{
		charPosDist{char3, position3, 3},
		charPosDist{char1, position1, 1},
		charPosDist{char2, position2, 2},
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
