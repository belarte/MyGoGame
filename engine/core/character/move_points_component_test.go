package character

import (
	"testing"
)

func TestMovePoints(t *testing.T) {
	handler := NewSimpleMovePointsComponent(5, 1)
	expected := 6.0
	result := handler.MovePoints()

	if result != expected {
		t.Errorf("Expected=%f, got %f", expected, result)
	}
}

func TestConsumeMovePoints(t *testing.T) {
	handler := SimpleMovePointsComponent{currentMP: 5}

	var expected = []bool{true, true, false}
	var results = []bool{handler.ConsumeMovePoints(2),
		handler.ConsumeMovePoints(2),
		handler.ConsumeMovePoints(2),
	}

	for i := range results {
		if expected[i] != results[i] {
			t.Errorf("%d> expected=%t, got=%t", i, expected[i], results[i])
		}
	}
}

func TestResetMovePoints(t *testing.T) {
	handler := NewSimpleMovePointsComponent(5, 1)

	var expected = []float64{6, 5, 1}
	var results = []float64{handler.currentMP, handler.baseMP, handler.bonus}

	for i := range results {
		if expected[i] != results[i] {
			t.Errorf("%d> expected=%f, got=%f", i, expected[i], results[i])
		}
	}
}
