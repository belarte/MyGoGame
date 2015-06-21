package engine

import (
	"testing"
)

func TestDistance(t *testing.T) {
	var expected = []int{
		distance(Coord{0, 0}, Coord{0, 0}),
		distance(Coord{0, 0}, Coord{0, 1}),
		distance(Coord{0, 0}, Coord{1, 0}),
		distance(Coord{0, 0}, Coord{1, 1}),
		distance(Coord{2, 2}, Coord{0, 1}),
		distance(Coord{2, 2}, Coord{0, 0}),
		distance(Coord{5, 5}, Coord{3, 1}),
	}

	var results = []int{0, 1, 1, 2, 3, 4, 6}

	for i, _ := range results {
		if results[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], results[i])
		}
	}
}
