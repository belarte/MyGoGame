package core

import (
	"testing"
)

func TestDistance(t *testing.T) {
	var expected = []float64{
		Distance(Coord{0, 0}, Coord{0, 0}),
		Distance(Coord{0, 0}, Coord{0, 1}),
		Distance(Coord{0, 0}, Coord{1, 0}),
		Distance(Coord{0, 0}, Coord{1, 1}),
		Distance(Coord{2, 2}, Coord{0, 1}),
		Distance(Coord{2, 2}, Coord{0, 0}),
		Distance(Coord{5, 5}, Coord{2, 1}),
	}

	var results = []float64{0., 1., 1., 1.414214, 2.236068, 2.828427, 5.}

	for i, _ := range results {
		if !CompareEpsilon(results[i], expected[i]) {
			t.Errorf("Expected %f, got %f", expected[i], results[i])
		}
	}
}
