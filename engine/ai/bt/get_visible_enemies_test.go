package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestGetVisibleEnemiesCheckConditionsNilLevel(t *testing.T) {
	char := &character.Actor{}
	context := &context{agent: char}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNilCharacter(t *testing.T) {
	lvl := level.New(utils.Coord{X: 0, Y: 0}, 0)
	context := &context{lvl: lvl}
	task := NewGetVisibleEnemies(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVisibleEnemiesCheckConditionsNotNilLevel(t *testing.T) {
	lvl := level.New(utils.Coord{X: 0, Y: 0}, 0)
	char := &character.Actor{}
	context := newContext(lvl, char)
	task := NewGetVisibleEnemies(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVisibleEnemiesPerformNoEnemiesOnLevel(t *testing.T) {
	lvl := level.New(utils.Coord{X: 0, Y: 0}, 2)
	char := &character.Actor{}
	lvl.AddCharacter(char, utils.Coord{X: 0, Y: 0}, 0)
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
	lvl := level.New(utils.Coord{X: 1, Y: 10}, 2)
	char1 := &character.Actor{
		StatsComponent:    &character.FakeStatsComponent{FakeVisibility: character.DefaultVisibility},
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.Coord{X: 0, Y: 0}},
	}
	char2 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.Coord{X: 0, Y: 9}},
	}
	lvl.AddCharacter(char1, utils.Coord{X: 0, Y: 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{X: 0, Y: 9}, 1)
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
	lvl := level.New(utils.Coord{X: 1, Y: 10}, 2)
	lvl.SetCell(utils.Coord{X: 0, Y: 1}, level.WallCell)
	char1 := &character.Actor{
		StatsComponent:    &character.FakeStatsComponent{FakeVisibility: character.DefaultVisibility},
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.Coord{X: 0, Y: 0}},
	}
	char2 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.Coord{X: 0, Y: 4}},
	}
	lvl.AddCharacter(char1, utils.Coord{X: 0, Y: 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{X: 0, Y: 4}, 1)
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
	lvl := level.New(utils.Coord{X: 1, Y: 10}, 2)
	char1 := &character.Actor{
		StatsComponent:    &character.FakeStatsComponent{FakeVisibility: character.DefaultVisibility},
		PositionComponent: &character.FakePositionComponent{},
	}
	char2 := &character.Actor{
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl.AddCharacter(char1, utils.Coord{X: 0, Y: 0}, 0)
	lvl.AddCharacter(char2, utils.Coord{X: 0, Y: 5}, 1)
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
