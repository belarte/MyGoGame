package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/action"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestMoveToDestinationConstruction(t *testing.T) {
	context := &context{
		lvl:             level.NewLevel(utils.Coord{1, 1}, 1),
		positionOfAgent: utils.Coord{0, 0},
		destination:     utils.Coord{0, 0},
	}

	task := NewMoveToDestination(context)

	if task == nil {
		t.Errorf("Construction failed: context=%+v", context)
	}
}

func TestMoveToDestinationCheckConditionsIsDoableFail(t *testing.T) {
	action := &action.Fake{FakeIsDoable: false}
	task := MoveToDestination{moveAction: action}

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, task=%+v", task)
	}
}

func TestMoveToDestinationCheckConditionsIsDoableSuccess(t *testing.T) {
	action := &action.Fake{FakeIsDoable: true}
	task := MoveToDestination{moveAction: action}

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true: task=%+v", task)
	}
}

func TestMoveToDestinationPerformActionPerformFail(t *testing.T) {
	action := &action.Fake{FakeIsDoable: false}
	task := MoveToDestination{moveAction: action}

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, task=%+v", task)
	}
}

func TestMoveToDestinationPerformActionPerformSuccess(t *testing.T) {
	action := &action.Fake{FakePerform: true}
	task := MoveToDestination{moveAction: action}

	if !task.Perform() {
		t.Errorf("CheckConditions should return true: task=%+v", task)
	}
}
