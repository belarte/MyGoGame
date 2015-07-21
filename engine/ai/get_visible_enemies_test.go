package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := &MockCharacter{}
	context := &context{agent: char}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 0)
	context := &context{level: level}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNotNilLevel(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 0)
	char := &MockCharacter{}
	context := newContext(level, char)
	task := NewGetVisibleEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesOnLevel(t *testing.T) {
	level := NewLevel(Coord{0, 0}, 2)
	char := &MockCharacter{}
	level.AddCharacter(char, Coord{0, 0}, 0)
	context := newContext(level, char)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesVisible(t *testing.T) {
	level := NewLevel(Coord{1, 10}, 2)
	char1 := &MockCharacter{}
	char2 := &MockCharacter{}
	level.AddCharacter(char1, Coord{0, 0}, 0)
	level.AddCharacter(char2, Coord{0, 9}, 1)
	context := newContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformVisibleEnemies(t *testing.T) {
	level := NewLevel(Coord{1, 10}, 2)
	char1 := &MockCharacter{VisibilityMock: DEFAULT_VISIBILITY}
	char2 := &MockCharacter{}
	level.AddCharacter(char1, Coord{0, 0}, 0)
	level.AddCharacter(char2, Coord{0, 5}, 1)
	context := newContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}
}
