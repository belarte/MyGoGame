package utils

import (
	"testing"
)

func TestBresenhamFromEqualsTo(t *testing.T) {
	result := Bresenham(Coord{0, 0}, Coord{0, 0})
	expectedSize := 0

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}
}

func TestBresenhamLineLeft(t *testing.T) {
	result := Bresenham(Coord{4, 0}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 0}, Coord{2, 0}, Coord{1, 0}, Coord{0, 0},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineRight(t *testing.T) {
	result := Bresenham(Coord{0, 0}, Coord{4, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 0}, Coord{2, 0}, Coord{3, 0}, Coord{4, 0},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineUp(t *testing.T) {
	result := Bresenham(Coord{0, 0}, Coord{0, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{0, 1}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineDown(t *testing.T) {
	result := Bresenham(Coord{0, 4}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{0, 3}, Coord{0, 2}, Coord{0, 1}, Coord{0, 0},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamBresenhamLineLeftUp(t *testing.T) {
	result := Bresenham(Coord{4, 0}, Coord{0, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 1}, Coord{2, 2}, Coord{1, 3}, Coord{0, 4},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineRightUp(t *testing.T) {
	result := Bresenham(Coord{0, 0}, Coord{4, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 1}, Coord{2, 2}, Coord{3, 3}, Coord{4, 4},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineLeftDown(t *testing.T) {
	result := Bresenham(Coord{4, 4}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 3}, Coord{2, 2}, Coord{1, 1}, Coord{0, 0},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamLineRightDown(t *testing.T) {
	result := Bresenham(Coord{0, 4}, Coord{4, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 3}, Coord{2, 2}, Coord{3, 1}, Coord{4, 0},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestBresenhamComplex(t *testing.T) {
	result := Bresenham(Coord{0, 0}, Coord{15, 10})
	expectedSize := 15

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 1}, Coord{2, 1}, Coord{3, 2}, Coord{4, 3}, Coord{5, 3},
		Coord{6, 4}, Coord{7, 5}, Coord{8, 5}, Coord{9, 6}, Coord{10, 7},
		Coord{11, 7}, Coord{12, 8}, Coord{13, 9}, Coord{14, 9}, Coord{15, 10},
	}

	for i, c := range result {
		if !EqualCoord(c, expected[i]) {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}
