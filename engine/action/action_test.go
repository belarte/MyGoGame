package action

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestMoveActionIsDoableEmptyPath(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &MockCharacter{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(level, char, &Path{})

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestMoveActionIsDoableNilPath(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &MockCharacter{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	action := NewMoveAction(level, char, nil)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestMoveActionIsDoableOK(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &MockCharacter{}
	level.AddCharacter(char, Coord{0, 0}, 0)

	var path Path
	path.Add(Coord{0, 1}, 1)
	action := NewMoveAction(level, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}
}

func TestMoveActionPerformOk(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &MockCharacter{MovePointsMock: 10}
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

func TestMoveActionPerformNotOk(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 1)
	char := &MockCharacter{MovePointsMock: 10}
	level.AddCharacter(char, Coord{0, 0}, 0)

	var path Path
	path.Add(Coord{0, 1}, 1)
	action := NewMoveAction(level, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}
}

func TestAttackActionIsDoable(t *testing.T) {
	t.Error("TODO")
}

func TestAttackActionPerform(t *testing.T) {
	t.Error("TODO")
}
