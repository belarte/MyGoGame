package action

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	. "github.com/belarte/MyGoGame/engine/utils"
)

func TestIsDoableEmptyPath(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Fake{}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(lvl, char, &level.Path{})

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableNilPath(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Fake{}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(lvl, char, nil)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoablePathDoesNotStartNextToAgent(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Fake{}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(Coord{0, 4}, 1)
	action := NewMoveAction(lvl, char, &path)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableOK(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Fake{}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(Coord{0, 1}, 1)
	action := NewMoveAction(lvl, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}
}

func TestPerformOk(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	dest := Coord{0, 1}
	char := &character.Fake{
		FakeMovePointsComponent: character.FakeMovePointsComponent{FakeMovePoints: 10, FakeConsumeMP: true},
		FakePositionComponent:   character.FakePositionComponent{FakePosition: dest},
	}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(dest, 1)
	action := NewMoveAction(lvl, char, &path)

	if !action.Perform() {
		t.Error("Move action should have performed.")
	}

	pos := char.Position()
	if pos != dest {
		t.Errorf("Desitnation not reached, expected %+v, is %+v.", dest, pos)
	}
}

func TestPerformNotEnoughMovePoints(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Fake{
		FakeMovePointsComponent: character.FakeMovePointsComponent{FakeMovePoints: 1},
	}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(Coord{0, 1}, 1)
	path.Add(Coord{0, 2}, 1)
	path.Add(Coord{0, 3}, 1)
	action := NewMoveAction(lvl, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}

	if action.Perform() {
		t.Errorf("Should not be able to perfrom: path=%+v, action=%+v", path, action)
	}

	pos := char.Position()
	dest := Coord{0, 3}
	if pos == dest {
		t.Errorf("Position of character is %+v, should be different than %+v", pos, dest)
	}
}

func TestPerformHasConsumedMovePoints(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	dest := Coord{0, 1}
	char := &character.Fake{
		FakeMovePointsComponent: character.FakeMovePointsComponent{FakeMovePoints: 10, FakeConsumeMP: true},
		FakePositionComponent:   character.FakePositionComponent{FakePosition: dest},
	}
	lvl.AddCharacter(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(dest, 1)
	action := NewMoveAction(lvl, char, &path)

	oldValue := char.MovePoints()
	if !action.Perform() {
		t.Error("Move action should have performed.")
	}

	newValue := char.MovePoints()
	if oldValue == newValue {
		t.Error("Move points should have been consumed.")
	}
}
