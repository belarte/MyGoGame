package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := &character.Fake{}
	context := &context{agent: char}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	lvl := level.New(utils.Coord{0, 0}, 0)
	context := &context{lvl: lvl}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNotNilLevel(t *testing.T) {
	lvl := level.New(utils.Coord{0, 0}, 0)
	char := &character.Fake{}
	context := newContext(lvl, char)
	task := NewGetVisibleEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesOnLevel(t *testing.T) {
	lvl := level.New(utils.Coord{0, 0}, 2)
	char := &character.Fake{}
	lvl.AddCharacter(char, utils.Coord{0, 0}, 0)
	context := newContext(lvl, char)
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

func TestGetVisibleEnemiesPerformNoEnemiesVisibleBecauseOfDistance(t *testing.T) {
	lvl := level.New(utils.Coord{1, 10}, 2)
	char1 := &character.Fake{FakeVisibility: character.DefaultVisibility}
	char2 := &character.Fake{}
	lvl.AddCharacter(char1, utils.Coord{0, 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{0, 9}, 1)
	context := newContext(lvl, char1)
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

func TestGetVisibleEnemiesPerformNoEnemiesVisibleBecauseOfWall(t *testing.T) {
	lvl := level.New(utils.Coord{1, 10}, 2)
	lvl.Map().SetCell(utils.Coord{0, 1}, level.WALL)
	char1 := &character.Fake{FakeVisibility: character.DefaultVisibility}
	char2 := &character.Fake{}
	lvl.AddCharacter(char1, utils.Coord{0, 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{0, 4}, 1)
	context := newContext(lvl, char1)
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
	lvl := level.New(utils.Coord{1, 10}, 2)
	char1 := &character.Fake{FakeVisibility: character.DefaultVisibility}
	char2 := &character.Fake{}
	lvl.AddCharacter(char1, utils.Coord{0, 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{0, 5}, 1)
	context := newContext(lvl, char1)
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
