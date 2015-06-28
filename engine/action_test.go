package engine

import (
	"testing"
)

func TestMoveActionIsDoableEmptyPath(t *testing.T) {
	level := NewLevel(Coord{5, 5}, 1)
	char := NewCharacter("Tester", 0, 0, 0, 0, 0)
	pos := Coord{0, 0}
	level.AddCharacter(char, pos, 0)
	action := NewMoveAction(level, char, pos)

	result := action.IsDoable()
	expected := false

	if result != expected {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func TestMoveActionIsDoableNotEnoughMP(t *testing.T) {
	level := NewLevel(Coord{5, 5}, 1)
	char := NewCharacter("Tester", 0, 0, 1, 0, 0)
	level.AddCharacter(char, Coord{0, 0}, 0)
	action := NewMoveAction(level, char, Coord{2, 2})

	result := action.IsDoable()
	expected := false

	if result != expected {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func TestMoveActionIsDoableOK(t *testing.T) {
	level := NewLevel(Coord{5, 5}, 1)
	char := NewCharacter("Tester", 0, 0, 10, 0, 0)
	level.AddCharacter(char, Coord{0, 0}, 0)
	action := NewMoveAction(level, char, Coord{2, 2})

	result := action.IsDoable()
	expected := true

	if result != expected {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func TestMoveActionPerform(t *testing.T) {
	t.Error("TODO")
}

func TestAttackActionIsDoable(t *testing.T) {
	t.Error("TODO")
}

func TestAttackActionPerform(t *testing.T) {
	t.Error("TODO")
}
