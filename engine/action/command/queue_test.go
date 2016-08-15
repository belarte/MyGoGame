package command

import "testing"

func TestSizeOfEmptyQueue(t *testing.T) {
	queue := NewQueue()

	expectedSize := 0
	size := queue.Size()

	if size != expectedSize {
		t.Errorf("Wrong size: expected %d, got %d", expectedSize, size)
	}
}

func TestSizeAfterAddCommandToQueue(t *testing.T) {
	queue := NewQueue()

	expectedSize := 3

	for i := 0; i < expectedSize; i++ {
		queue.Add(&Fake{})
	}

	size := queue.Size()

	if size != expectedSize {
		t.Errorf("Wrong size: expected %d, got %d", expectedSize, size)
	}
}

func TestSizeAfterFlush(t *testing.T) {
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Add(&Fake{})
	}
	queue.Flush()

	expectedSize := 0
	size := queue.Size()

	if size != expectedSize {
		t.Errorf("Wrong size: expected %d, got %d", expectedSize, size)
	}
}

func TestProcessConsumeCommands(t *testing.T) {
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Add(&Fake{})
	}

	for i := 0; i < 5; i++ {
		queue.ProcessNext()
	}

	expectedSize := 0
	size := queue.Size()

	if size != expectedSize {
		t.Errorf("Wrong size: expected %d, got %d", expectedSize, size)
	}
}
