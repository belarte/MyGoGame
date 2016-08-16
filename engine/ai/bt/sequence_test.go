package bt

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
	subTask := &Fake{}
	task := NewSequence(nil)
	task.Add(subTask)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskCheckConditionsFailed(t *testing.T) {
	subTask := &Fake{FakeCheckConditions: false}
	task := NewSequence(nil)
	task.Add(subTask)

	if task.Perform() == success {
		t.Errorf("CheckConditions should return false,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskPerformFailed(t *testing.T) {
	subTask := &Fake{FakeCheckConditions: true, FakePerform: failure}
	task := NewSequence(nil)
	task.Add(subTask)

	if task.Perform() == success {
		t.Errorf("CheckConditions should return false,\ntasks=%+v", task.tasks)
	}
}

func TestSequencePerformSubTaskPerformSucceeded(t *testing.T) {
	subTask := &Fake{FakeCheckConditions: true, FakePerform: success}
	task := NewSequence(nil)
	task.Add(subTask)
	task.Add(subTask)

	if task.Perform() == failure {
		t.Errorf("CheckConditions should return true,\ntasks=%+v", task.tasks)
	}
}
