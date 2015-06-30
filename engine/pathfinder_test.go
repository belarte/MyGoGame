package engine

import (
	"testing"
)

func TestPathToSelfTesting(t *testing.T) {
	level := NewLevel(NewCoord(1, 1), 0)
	finder := NewPathFinder(level)

	shortest := finder.ShortestPath(Coord{0, 0}, Coord{0, 0})

	if shortest.size() != 0 {
		t.Error("shortest.size(): expected 0, got ", shortest.size())
	}

	if shortest.cost() != 0 {
		t.Error("shortest.cost(): expected 0, got ", shortest.cost())
	}
}

func TestPathInLineTesting(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 0)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 4
	expectedCost := 4.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathInColumnTesting(t *testing.T) {
	level := NewLevel(Coord{5, 1}, 0)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{4, 0})
	shortest2 := finder.ShortestPath(Coord{4, 0}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 4
	expectedCost := 4.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathWithObstacleTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.maps.SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 6
	expectedCost := 6.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesOverTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.maps.SetCell(Coord{0, 1}, DIFFICULT)
	level.maps.SetCell(Coord{0, 2}, DIFFICULT)
	level.maps.SetCell(Coord{0, 3}, DIFFICULT)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 6
	expectedCost := 6.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesCrossTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.maps.SetCell(Coord{0, 1}, DIFFICULT)
	level.maps.SetCell(Coord{0, 2}, DIFFICULT)
	level.maps.SetCell(Coord{0, 3}, DIFFICULT)
	level.maps.SetCell(Coord{1, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 4
	expectedCost := 7.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathWithProtagonistTesting(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{2, 5}, 1)
	level.AddCharacter(char, Coord{0, 2}, 0)

	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 6
	expectedCost := 6.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestPathNotPossibleTesting(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 0)
	level.maps.SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.size(),
		shortest2.size(),
	}

	var costs = []float64{
		shortest1.cost(),
		shortest2.cost(),
	}

	expectedSize := 0
	expectedCost := 0.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if res != expectedCost {
			t.Errorf("Expected %d, got %d", expectedCost, res)
		}
	}
}

func TestTIsInList(t *testing.T) {
	fakeNode := node{0, 0, 0, Coord{0, 0}}
	list := nodeList{Coord{0, 0}: fakeNode,
		Coord{2, 1}: fakeNode,
		Coord{5, 6}: fakeNode,
	}

	finder := NewPathFinder(nil)

	results := []bool{finder.isInList(Coord{0, 0}, list),
		finder.isInList(Coord{1, 1}, list),
		finder.isInList(Coord{2, 1}, list),
		finder.isInList(Coord{1, 2}, list),
		finder.isInList(Coord{5, 6}, list),
		finder.isInList(Coord{6, 7}, list),
	}

	expected := []bool{true, false, true, false, true, false}

	for i, _ := range results {
		if expected[i] != results[i] {
			t.Errorf("Expected %t, got %t", expected[i], results[i])
		}
	}
}

func TestGetAdjacentCells(t *testing.T) {
	level := NewLevel(Coord{3, 3}, 0)
	finder := NewPathFinder(level)

	results := [][]Coord{finder.getAdjacentCells(Coord{0, 0}),
		finder.getAdjacentCells(Coord{0, 1}),
		finder.getAdjacentCells(Coord{0, 2}),
		finder.getAdjacentCells(Coord{1, 0}),
		finder.getAdjacentCells(Coord{1, 1}),
		finder.getAdjacentCells(Coord{1, 2}),
		finder.getAdjacentCells(Coord{2, 0}),
		finder.getAdjacentCells(Coord{2, 1}),
		finder.getAdjacentCells(Coord{2, 2}),
	}

	expected := []int{2, 3, 2, 3, 4, 3, 2, 3, 2}

	for i, _ := range results {
		if expected[i] != len(results[i]) {
			t.Errorf("Expected %d, got %d", expected[i], results[i])
		}
	}
}
