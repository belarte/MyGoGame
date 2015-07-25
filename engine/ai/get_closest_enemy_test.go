package ai

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
	char := &character.Mock{}
	position := utils.Coord{0, 1}
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
	char1 := &character.Mock{}
	char2 := &character.Mock{}
	char3 := &character.Mock{}
	position1 := utils.Coord{0, 1}
	position2 := utils.Coord{0, 2}
	position3 := utils.Coord{0, 3}

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
