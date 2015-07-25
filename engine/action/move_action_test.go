package action

import (
	"testing"

	. "github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
	. "github.com/belarte/MyGoGame/engine/utils"
)

func TestIsDoableEmptyPath(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(level, char, &Path{})

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableNilPath(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(level, char, nil)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoablePathDoesNotStartNextToAgent(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	var path Path
	path.Add(Coord{0, 4}, 1)
	action := NewMoveAction(level, char, &path)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableOK(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	var path Path
	path.Add(Coord{0, 1}, 1)
	action := NewMoveAction(level, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}
}

func TestPerformOk(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{MovePointsMock: 10, ConsumeMPMock: true}
	level.AddCharacter(char, Coord{0, 0}, 0)
	dest := Coord{0, 1}

	var path Path
	path.Add(dest, 1)
	action := NewMoveAction(level, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}

	if !action.Perform() {
		t.Error("Move action should have performed.")
	}

	pos := level.PositionOf(char)
	if pos != dest {
		t.Errorf("Desitnation not reached, expected %+v, is %+v.", dest, pos)
	}
}

func TestPerformNotEnoughMovePoints(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{MovePointsMock: 1}
	level.AddCharacter(char, Coord{0, 0}, 0)

	var path Path
	path.Add(Coord{0, 1}, 1)
	path.Add(Coord{0, 2}, 1)
	path.Add(Coord{0, 3}, 1)
	action := NewMoveAction(level, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}

	if action.Perform() {
		t.Errorf("Should not be able to perfrom: path=%+v, action=%+v", path, action)
	}

	pos := level.PositionOf(char)
	dest := Coord{0, 3}
	if pos == dest {
		t.Errorf("Position of character is %+v, should be different than %+v", pos, dest)
	}
}

func TestPerformHasConsumedMovePoints(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &character.Mock{MovePointsMock: 10, ConsumeMPMock: true}
	level.AddCharacter(char, Coord{0, 0}, 0)
	dest := Coord{0, 1}

	var path Path
	path.Add(dest, 1)
	action := NewMoveAction(level, char, &path)

	oldValue := char.MovePoints()
	if !action.Perform() {
		t.Error("Move action should have performed.")
	}

	newValue := char.MovePoints()
	if oldValue == newValue {
		t.Error("Move points should have been consumed.")
	}
}
