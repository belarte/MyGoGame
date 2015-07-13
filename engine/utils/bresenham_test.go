package utils

import (
	"testing"
)

func TestLineFromEqualsTo(t *testing.T) {
	result := Line(Coord{0, 0}, Coord{0, 0})
	expectedSize := 0

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}
}

func TestLineLineLeft(t *testing.T) {
	result := Line(Coord{4, 0}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 0}, Coord{2, 0}, Coord{1, 0}, Coord{0, 0},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineRight(t *testing.T) {
	result := Line(Coord{0, 0}, Coord{4, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 0}, Coord{2, 0}, Coord{3, 0}, Coord{4, 0},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineUp(t *testing.T) {
	result := Line(Coord{0, 0}, Coord{0, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{0, 1}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineDown(t *testing.T) {
	result := Line(Coord{0, 4}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{0, 3}, Coord{0, 2}, Coord{0, 1}, Coord{0, 0},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineLineLeftUp(t *testing.T) {
	result := Line(Coord{4, 0}, Coord{0, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 1}, Coord{2, 2}, Coord{1, 3}, Coord{0, 4},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineRightUp(t *testing.T) {
	result := Line(Coord{0, 0}, Coord{4, 4})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 1}, Coord{2, 2}, Coord{3, 3}, Coord{4, 4},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineLeftDown(t *testing.T) {
	result := Line(Coord{4, 4}, Coord{0, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{3, 3}, Coord{2, 2}, Coord{1, 1}, Coord{0, 0},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineLineRightDown(t *testing.T) {
	result := Line(Coord{0, 4}, Coord{4, 0})
	expectedSize := 4

	if len(result) != expectedSize {
		t.Fatalf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = []Coord{
		Coord{1, 3}, Coord{2, 2}, Coord{3, 1}, Coord{4, 0},
	}

	for i, c := range result {
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestLineComplex(t *testing.T) {
	result := Line(Coord{0, 0}, Coord{15, 10})
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
		if c != expected[i] {
			t.Errorf("%d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestCircle(t *testing.T) {
	result := Circle(Coord{9, 9}, 7)
	expectedSize := 40

	if len(result) != expectedSize {
		t.Errorf("Expected size=%d, got %d", expectedSize, len(result))
	}

	var expected = map[Coord]bool{
		Coord{16, 9}: true, Coord{16, 10}: true,
		Coord{16, 11}: true, Coord{15, 12}: true,
		Coord{15, 13}: true, Coord{14, 14}: true,
		Coord{13, 15}: true, Coord{12, 15}: true,
		Coord{11, 16}: true, Coord{10, 16}: true,
		Coord{9, 16}: true, Coord{8, 16}: true,
		Coord{7, 16}: true, Coord{6, 15}: true,
		Coord{5, 15}: true, Coord{4, 14}: true,
		Coord{3, 13}: true, Coord{3, 12}: true,
		Coord{2, 11}: true, Coord{2, 10}: true,
		Coord{2, 9}: true, Coord{2, 8}: true,
		Coord{2, 7}: true, Coord{3, 6}: true,
		Coord{3, 5}: true, Coord{4, 4}: true,
		Coord{5, 3}: true, Coord{6, 3}: true,
		Coord{7, 2}: true, Coord{8, 2}: true,
		Coord{9, 2}: true, Coord{10, 2}: true,
		Coord{11, 2}: true, Coord{12, 3}: true,
		Coord{13, 3}: true, Coord{14, 4}: true,
		Coord{15, 5}: true, Coord{15, 6}: true,
		Coord{16, 7}: true, Coord{16, 8}: true,
	}

	for i, c := range result {
		if !expected[c] {
			t.Errorf("%d: %+v is not an expected result", i, c)
		}
	}
}
