package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	context := &Context{nil, char, Coord{0, 0}, nil, nil, nil}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 0)
	context := &Context{level, nil, Coord{0, 0}, nil, nil, nil}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNotNilLevel(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 0)
	char := NewCharacter("", 0, 0, 0, 0, 0)
	context := NewContext(level, char)
	task := NewGetVisibleEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesOnLevel(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 2)
	char := NewCharacter("", 0, 0, 0, 0, 0)
	level.AddCharacter(char, Coord{0, 0}, 0)
	context := NewContext(level, char)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesVisible(t *testing.T) {
	level := NewLevel(Coord{1, 10}, 2)
	char1 := NewCharacter("", 0, 0, 0, 0, 0)
	char2 := NewCharacter("", 0, 0, 0, 0, 0)
	level.AddCharacter(char1, Coord{0, 0}, 0)
	level.AddCharacter(char2, Coord{0, 9}, 1)
	context := NewContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformVisibleEnemies(t *testing.T) {
	level := NewLevel(Coord{1, 10}, 2)
	char1 := NewCharacter("", 0, 0, 0, 0, 0)
	char2 := NewCharacter("", 0, 0, 0, 0, 0)
	level.AddCharacter(char1, Coord{0, 0}, 0)
	level.AddCharacter(char2, Coord{0, 5}, 1)
	context := NewContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}
}

func TestGetClosestEnemyCheckConditionsNoVisibleEnemies(t *testing.T) {
	context := &Context{nil, nil, Coord{0, 0}, make([]charPosDist, 0, 0), nil, nil}
	task := NewGetClosestEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetClosestEnemyCheckConditionsVisibleEnemies(t *testing.T) {
	context := &Context{nil, nil, Coord{0, 0}, make([]charPosDist, 1, 1), nil, nil}
	task := NewGetClosestEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetClosestEnemyPerformOneEnemy(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	position := Coord{0, 1}
	opponents := []charPosDist{charPosDist{char, &position, 1}}
	context := &Context{nil, nil, Coord{0, 0}, opponents, nil, nil}
	task := NewGetClosestEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}

	if context.closestEnemy != char {
		t.Errorf("Expected %+v, got %+v", char, context.closestEnemy)
	}

	if !EqualCoord(*(context.closestEnemyPosition), position) {
		t.Errorf("Expected %+v, got %+v", position, context.closestEnemyPosition)
	}
}

func TestGetClosestEnemyPerformThreeEnemies(t *testing.T) {
	char1 := NewCharacter("char1", 0, 0, 0, 0, 0)
	char2 := NewCharacter("char2", 0, 0, 0, 0, 0)
	char3 := NewCharacter("char3", 0, 0, 0, 0, 0)
	position1 := Coord{0, 1}
	position2 := Coord{0, 2}
	position3 := Coord{0, 3}

	opponents := []charPosDist{
		charPosDist{char3, &position3, 3},
		charPosDist{char1, &position1, 1},
		charPosDist{char2, &position2, 2},
	}
	context := &Context{nil, nil, Coord{0, 0}, opponents, nil, nil}
	task := NewGetClosestEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}

	if context.closestEnemy != char1 {
		t.Errorf("Expected %v, got %+v", char1, context.closestEnemy)
	}

	if !EqualCoord(*(context.closestEnemyPosition), position1) {
		t.Errorf("Expected %+v, got %+v", position1, context.closestEnemyPosition)
	}
}
