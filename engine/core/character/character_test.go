package character

import (
	"testing"
)

/*
type movePointsHandler struct {
	currentMP, baseMP, bonus float64
}

func (handler *movePointsHandler) MovePoints() float64 {
	return handler.currentMP
}

func (handler *movePointsHandler) ConsumeMovePoints(points float64) bool {
	if points > handler.currentMP {
		return false
	}

	handler.currentMP -= points
	return true
}

func (handler *movePointsHandler) Reset() {
	handler.currentMP = handler.baseMP + handler.bonus
}
*/

func TestMPHandlerMovePoints(t *testing.T) {
	handler := newMovePointsHandler(5, 1)
	expected := 6.0
	result := handler.MovePoints()

	if result != expected {
		t.Errorf("Expected=%f, got %f", expected, result)
	}
}

func TestMPHandlerConsume(t *testing.T) {
	handler := movePointsHandler{currentMP: 5}

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

func TestMPHandlerReset(t *testing.T) {
	handler := newMovePointsHandler(5, 1)

	var expected = []float64{6, 5, 1}
	var results = []float64{handler.currentMP, handler.baseMP, handler.bonus}

	for i := range results {
		if expected[i] != results[i] {
			t.Errorf("%d> expected=%f, got=%f", i, expected[i], results[i])
		}
	}
}
