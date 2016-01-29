package ai

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := &character.Mock{}
	context := &context{agent: char}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	level := core.NewLevel(utils.Coord{0, 0}, 0)
	context := &context{level: level}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNotNilLevel(t *testing.T) {
	level := core.NewLevel(utils.Coord{0, 0}, 0)
	char := &character.Mock{}
	context := newContext(level, char)
	task := NewGetVisibleEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesOnLevel(t *testing.T) {
	level := core.NewLevel(utils.Coord{0, 0}, 2)
	char := &character.Mock{}
	level.AddCharacter(char, utils.Coord{0, 0}, 0)
	context := newContext(level, char)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}

	expectedVisible := 0
	resultVisible := len(context.visibleEnemies)

	if resultVisible != expectedVisible {
		t.Errorf("Wrong number of visible enemies\nresult=%d, expected=%d", resultVisible, expectedVisible)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesVisible(t *testing.T) {
	level := core.NewLevel(utils.Coord{1, 10}, 2)
	char1 := &character.Mock{}
	char2 := &character.Mock{}
	level.AddCharacter(char1, utils.Coord{0, 0}, 0)
	level.AddCharacter(char2, utils.Coord{0, 9}, 1)
	context := newContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if task.Perform() {
		t.Errorf("Perform should be false, context=%+v", context)
	}

	expectedVisible := 0
	resultVisible := len(context.visibleEnemies)

	if resultVisible != expectedVisible {
		t.Errorf("Wrong number of visible enemies\nresult=%d, expected=%d", resultVisible, expectedVisible)
	}
}

func TestGetVisibleEnemiesPerformVisibleEnemies(t *testing.T) {
	level := core.NewLevel(utils.Coord{1, 10}, 2)
	char1 := &character.Mock{VisibilityMock: character.DEFAULT_VISIBILITY}
	char2 := &character.Mock{}
	level.AddCharacter(char1, utils.Coord{0, 0}, 0)
	level.AddCharacter(char2, utils.Coord{0, 5}, 1)
	context := newContext(level, char1)
	task := NewGetVisibleEnemies(context)

	if !task.Perform() {
		t.Errorf("Perform should be true, context=%+v", context)
	}

	expectedVisible := 1
	resultVisible := len(context.visibleEnemies)

	if resultVisible != expectedVisible {
		t.Errorf("Wrong number of visible enemies\nresult=%d, expected=%d", resultVisible, expectedVisible)
	}
}
