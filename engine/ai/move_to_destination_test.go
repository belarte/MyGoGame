package ai

import (
	. "github.com/belarte/MyGoGame/engine/action"
	"testing"
)

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
