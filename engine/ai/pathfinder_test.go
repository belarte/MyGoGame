package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestPathToSelfTesting(t *testing.T) {
	level := NewLevel(NewCoord(1, 1), 0)
	finder := NewPathFinder(level)

	shortest := finder.ShortestPath(Coord{0, 0}, Coord{0, 0})

	if shortest.Size() != 0 {
		t.Error("shortest.Size(): expected 0, got ", shortest.Size())
	}

	if !CompareEpsilon(shortest.Cost(), 0) {
		t.Error("shortest.Cost(): expected 0, got ", shortest.Cost())
	}
}

func TestPathInLineTesting(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 0)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 4.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathInColumnTesting(t *testing.T) {
	level := NewLevel(Coord{5, 1}, 0)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{4, 0})
	shortest2 := finder.ShortestPath(Coord{4, 0}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 4.0

	for i, res := range sizes {
		if res != expectedSize {
			t.Errorf("%d: Expected size %d, got %d", i, expectedSize, res)
		}
	}

	for i, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("%d: Expected cost %f, got %f", i, expectedCost, res)
		}
	}
}

func TestPathWithObstacleTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.Map().SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 4.828427

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesOverTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.Map().SetCell(Coord{0, 1}, DIFFICULT)
	level.Map().SetCell(Coord{0, 2}, DIFFICULT)
	level.Map().SetCell(Coord{0, 3}, DIFFICULT)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 4.828427

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesCrossTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5}, 0)
	level.Map().SetCell(Coord{0, 1}, DIFFICULT)
	level.Map().SetCell(Coord{0, 2}, DIFFICULT)
	level.Map().SetCell(Coord{0, 3}, DIFFICULT)
	level.Map().SetCell(Coord{1, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 7.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
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
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 4
	expectedCost := 4.828427

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathNotPossibleTesting(t *testing.T) {
	level := NewLevel(Coord{1, 5}, 0)
	level.Map().SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var sizes = []int{
		shortest1.Size(),
		shortest2.Size(),
	}

	var costs = []float64{
		shortest1.Cost(),
		shortest2.Cost(),
	}

	expectedSize := 0
	expectedCost := 0.0

	for _, res := range sizes {
		if res != expectedSize {
			t.Errorf("Expected %d, got %d", expectedSize, res)
		}
	}

	for _, res := range costs {
		if !CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
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

	expected := []int{3, 5, 3, 5, 8, 5, 3, 5, 3}

	for i, _ := range results {
		if expected[i] != len(results[i]) {
			t.Errorf("%d", results[i])
			t.Errorf("%d: Expected %d, got %d", i, expected[i], len(results[i]))
		}
	}
}
