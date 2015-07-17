package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	context := &Context{agent: char}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 0)
	context := &Context{level: level}
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
	context := &Context{}
	task := NewGetClosestEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetClosestEnemyCheckConditionsVisibleEnemies(t *testing.T) {
	context := &Context{visibleEnemies: make([]charPosDist, 1, 1)}
	task := NewGetClosestEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetClosestEnemyPerformOneEnemy(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	position := Coord{0, 1}
	opponents := []charPosDist{charPosDist{char, position, 1}}
	context := &Context{visibleEnemies: opponents}
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
	char1 := NewCharacter("char1", 0, 0, 0, 0, 0)
	char2 := NewCharacter("char2", 0, 0, 0, 0, 0)
	char3 := NewCharacter("char3", 0, 0, 0, 0, 0)
	position1 := Coord{0, 1}
	position2 := Coord{0, 2}
	position3 := Coord{0, 3}

	opponents := []charPosDist{
		charPosDist{char3, position3, 3},
		charPosDist{char1, position1, 1},
		charPosDist{char2, position2, 2},
	}
	context := &Context{visibleEnemies: opponents}
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

func TestGetVantagePointCheckConditionsNoClosestEnemy(t *testing.T) {
	context := &Context{closestEnemyPosition: Coord{-1, -1}}
	task := NewGetVantagePoint(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVantagePointCheckConditionsClosestEnemy(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{0, 0}, 0)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 0}}
	task := NewGetVantagePoint(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVantagePointPerformNoObstacle(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{1, 5}, 1)
	level.AddCharacter(agent, Coord{0, 0}, 0)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := Coord{0, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacle(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{2, 5}, 1)
	level.AddCharacter(agent, Coord{0, 0}, 0)
	level.Map().SetCell(Coord{0, 3}, WALL)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := Coord{1, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}
