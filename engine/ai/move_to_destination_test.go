package ai

import (
	. "github.com/belarte/MyGoGame/engine/action"
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestMoveToDestinationConstruction(t *testing.T) {
	context := &context{
		level:           NewLevel(Coord{1, 1}, 1),
		positionOfAgent: Coord{0, 0},
		destination:     Coord{0, 0},
	}

	task := NewMoveToDestination(context)

	if task == nil {
		t.Errorf("Construction failed: context=%+v", context)
	}
}

func TestMoveToDestinationCheckConditionsIsDoableFail(t *testing.T) {
	action := &MockAction{IsDoableMock: false}
	task := MoveToDestination{moveAction: action}

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, task=%+v", task)
	}
}

func TestMoveToDestinationCheckConditionsIsDoableSuccess(t *testing.T) {
	action := &MockAction{IsDoableMock: true}
	task := MoveToDestination{moveAction: action}

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true: task=%+v", task)
	}
}

func TestMoveToDestinationPerformActionPerformFail(t *testing.T) {
	action := &MockAction{IsDoableMock: false}
	task := MoveToDestination{moveAction: action}

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, task=%+v", task)
	}
}

func TestMoveToDestinationPerformActionPerformSuccess(t *testing.T) {
	action := &MockAction{PerformMock: true}
	task := MoveToDestination{moveAction: action}

	if !task.Perform() {
		t.Errorf("CheckConditions should return true: task=%+v", task)
	}
}
