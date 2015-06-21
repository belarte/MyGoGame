package engine

import (
	"testing"
)

func TestPathToSelfTesting(t *testing.T) {
	level := NewLevel(NewCoord(1, 1))
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
	level := NewLevel(Coord{1, 5})
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 4

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
		}
	}
}

func TestPathInColumnTesting(t *testing.T) {
	level := NewLevel(Coord{5, 1})
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{4, 0})
	shortest2 := finder.ShortestPath(Coord{4, 0}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 4

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
		}
	}
}

func TestPathWithObstacleTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5})
	level.maps.SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 6

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
		}
	}
}

func TestPathWithDifficultiesOverTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5})
	level.maps.SetCell(Coord{0, 1}, DIFFICULT)
	level.maps.SetCell(Coord{0, 2}, DIFFICULT)
	level.maps.SetCell(Coord{0, 3}, DIFFICULT)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 6

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
		}
	}
}

func TestPathWithDifficultiesCrossTesting(t *testing.T) {
	level := NewLevel(Coord{2, 5})
	level.maps.SetCell(Coord{0, 1}, DIFFICULT)
	level.maps.SetCell(Coord{0, 2}, DIFFICULT)
	level.maps.SetCell(Coord{0, 3}, DIFFICULT)
	level.maps.SetCell(Coord{1, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	var expected = []int{4, 7, 4, 7}

	for i, _ := range results {
		if results[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], results[i])
		}
	}
}

func TestPathWithProtagonistTesting(t *testing.T) {
	char := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{2, 5})
	level.AddCharacter(char, Coord{0, 2})

	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 6

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
		}
	}
}

func TestPathNotPossibleTesting(t *testing.T) {
	level := NewLevel(Coord{1, 5})
	level.maps.SetCell(Coord{0, 2}, WALL)
	finder := NewPathFinder(level)

	shortest1 := finder.ShortestPath(Coord{0, 0}, Coord{0, 4})
	shortest2 := finder.ShortestPath(Coord{0, 4}, Coord{0, 0})

	var results = []int{
		shortest1.size(),
		shortest1.cost(),
		shortest2.size(),
		shortest2.cost(),
	}

	expected := 0

	for _, res := range results {
		if res != expected {
			t.Errorf("Expected %d, got %d", expected, res)
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
	level := NewLevel(Coord{3, 3})
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
