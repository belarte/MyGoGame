package utils

import (
	"testing"
)

func TestDistance(t *testing.T) {
	var results = []float64{
		Distance(Coord{0, 0}, Coord{0, 0}),
		Distance(Coord{0, 0}, Coord{0, 1}),
		Distance(Coord{0, 0}, Coord{1, 0}),
		Distance(Coord{0, 0}, Coord{1, 1}),
		Distance(Coord{2, 2}, Coord{0, 1}),
		Distance(Coord{2, 2}, Coord{0, 0}),
		Distance(Coord{5, 5}, Coord{2, 1}),
	}

	var expected = []float64{0., 1., 1., 1.414214, 2.236068, 2.828427, 5.}

	for i := range results {
		if !CompareEpsilon(results[i], expected[i]) {
			t.Errorf("Expected %f, got %f", expected[i], results[i])
		}
	}
}

func TestAreAdjacent(t *testing.T) {
	var results = []bool{
		AreAdjacent(Coord{0, 0}, Coord{0, 0}),
		AreAdjacent(Coord{0, 0}, Coord{0, 1}),
		AreAdjacent(Coord{0, 0}, Coord{1, 0}),
		AreAdjacent(Coord{0, 0}, Coord{1, 1}),
		AreAdjacent(Coord{2, 2}, Coord{0, 1}),
		AreAdjacent(Coord{2, 2}, Coord{0, 0}),
		AreAdjacent(Coord{5, 5}, Coord{2, 1}),
	}

	var expected = []bool{false, true, true, true, false, false, false}

	for i := range results {
		if results[i] != expected[i] {
			t.Errorf("%d> Expected %t, got %t", i, expected[i], results[i])
		}
	}
}
