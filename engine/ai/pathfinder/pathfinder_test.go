package pathfinder

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestPathToSelfTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 1, Y: 1}, 0)
	finder := New(lvl)

	shortest := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 0})

	if shortest.Size() != 0 {
		t.Error("shortest.Size(): expected 0, got ", shortest.Size())
	}

	if !utils.CompareEpsilon(shortest.Cost(), 0) {
		t.Error("shortest.Cost(): expected 0, got ", shortest.Cost())
	}
}

func TestPathInLineTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 1, Y: 5}, 0)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathInColumnTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 5, Y: 1}, 0)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 4, Y: 0})
	shortest2 := finder.ShortestPath(utils.Coord{X: 4, Y: 0}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("%d: Expected cost %f, got %f", i, expectedCost, res)
		}
	}
}

func TestPathWithObstacleTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 2, Y: 5}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 2}, level.WallCell)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesOverTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 2, Y: 5}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 1}, level.DifficultCell)
	lvl.SetCell(utils.Coord{X: 0, Y: 2}, level.DifficultCell)
	lvl.SetCell(utils.Coord{X: 0, Y: 3}, level.DifficultCell)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathWithDifficultiesCrossTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 2, Y: 5}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 1}, level.DifficultCell)
	lvl.SetCell(utils.Coord{X: 0, Y: 2}, level.DifficultCell)
	lvl.SetCell(utils.Coord{X: 0, Y: 3}, level.DifficultCell)
	lvl.SetCell(utils.Coord{X: 1, Y: 2}, level.WallCell)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathWithProtagonistTesting(t *testing.T) {
	char := &character.Fake{}
	lvl := level.New(utils.Coord{X: 2, Y: 5}, 1)
	lvl.AddCharacter(char, utils.Coord{X: 0, Y: 2}, 0)

	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestPathNotPossibleTesting(t *testing.T) {
	lvl := level.New(utils.Coord{X: 1, Y: 5}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 2}, level.WallCell)
	finder := New(lvl)

	shortest1 := finder.ShortestPath(utils.Coord{X: 0, Y: 0}, utils.Coord{X: 0, Y: 4})
	shortest2 := finder.ShortestPath(utils.Coord{X: 0, Y: 4}, utils.Coord{X: 0, Y: 0})

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
		if !utils.CompareEpsilon(res, expectedCost) {
			t.Errorf("Expected %f, got %f", expectedCost, res)
		}
	}
}

func TestIsInList(t *testing.T) {
	fakeNode := node{0, 0, 0, utils.Coord{X: 0, Y: 0}}
	list := nodeList{utils.Coord{X: 0, Y: 0}: fakeNode,
		utils.Coord{X: 2, Y: 1}: fakeNode,
		utils.Coord{X: 5, Y: 6}: fakeNode,
	}

	finder := New(nil)

	results := []bool{finder.isInList(utils.Coord{X: 0, Y: 0}, list),
		finder.isInList(utils.Coord{X: 1, Y: 1}, list),
		finder.isInList(utils.Coord{X: 2, Y: 1}, list),
		finder.isInList(utils.Coord{X: 1, Y: 2}, list),
		finder.isInList(utils.Coord{X: 5, Y: 6}, list),
		finder.isInList(utils.Coord{X: 6, Y: 7}, list),
	}

	expected := []bool{true, false, true, false, true, false}

	for i := range results {
		if expected[i] != results[i] {
			t.Errorf("Expected %t, got %t", expected[i], results[i])
		}
	}
}

func TestGetAdjacentCells(t *testing.T) {
	lvl := level.New(utils.Coord{X: 3, Y: 3}, 0)
	finder := New(lvl)

	results := [][]utils.Coord{finder.getAdjacentCells(utils.Coord{X: 0, Y: 0}),
		finder.getAdjacentCells(utils.Coord{X: 0, Y: 1}),
		finder.getAdjacentCells(utils.Coord{X: 0, Y: 2}),
		finder.getAdjacentCells(utils.Coord{X: 1, Y: 0}),
		finder.getAdjacentCells(utils.Coord{X: 1, Y: 1}),
		finder.getAdjacentCells(utils.Coord{X: 1, Y: 2}),
		finder.getAdjacentCells(utils.Coord{X: 2, Y: 0}),
		finder.getAdjacentCells(utils.Coord{X: 2, Y: 1}),
		finder.getAdjacentCells(utils.Coord{X: 2, Y: 2}),
	}

	expected := []int{3, 5, 3, 5, 8, 5, 3, 5, 3}

	for i := range results {
		if expected[i] != len(results[i]) {
			t.Errorf("%d", results[i])
			t.Errorf("%d: Expected %d, got %d", i, expected[i], len(results[i]))
		}
	}
}
