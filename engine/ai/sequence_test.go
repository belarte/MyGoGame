package ai

import (
	"testing"
)

func TestSequenceCheckConditionsNoTasks(t *testing.T) {
	task := NewSequence(nil)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false,\ntasks=%+v", task.tasks)
	}
}

func TestSequenceCheckConditionsTasksListNotEmpty(t *testing.T) {
	subTask := &MockTask{}
	task := NewSequence(nil)
	task.Add(subTask)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskCheckConditionsFailed(t *testing.T) {
	subTask := &MockTask{CheckConditionsMock: false}
	task := NewSequence(nil)
	task.Add(subTask)

	if task.Perform() {
		t.Errorf("CheckConditions should return false,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskPerformFailed(t *testing.T) {
	subTask := &MockTask{CheckConditionsMock: true, PerformMock: false}
	task := NewSequence(nil)
	task.Add(subTask)

	if task.Perform() {
		t.Errorf("CheckConditions should return false,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskPerformSucceeded(t *testing.T) {
	subTask := &MockTask{CheckConditionsMock: true, PerformMock: true}
	task := NewSequence(nil)
	task.Add(subTask)
	task.Add(subTask)

	if !task.Perform() {
		t.Errorf("CheckConditions should return true,\ntasks=%+v", task.tasks)
	}
}
