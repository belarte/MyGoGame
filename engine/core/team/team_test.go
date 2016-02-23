package team

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestEmptyTeamIsEmpty(t *testing.T) {
	team := New()

	size := team.CharactersCount()
	expectedSize := 0

	if size != expectedSize {
		t.Errorf("Expected=%d, Result=%d", expectedSize, size)
	}
}

func TestFullTeamIsFull(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddCharacter(&character.Fake{}, utils.NilCoord)
	}

	if !team.IsFull() {
		t.Error("team should be full")
	}

	if team.CharactersCount() != MaxPlayersByTeam {
		t.Errorf("team should have %d characters, but has %d", MaxPlayersByTeam, team.CharactersCount())
	}
}

func TestEmptyTeamReturnsEmptyListOfCharacters(t *testing.T) {
	team := New()
	list := team.GetCharacters()

	if len(list) != 0 {
		t.Errorf("List should be empty, but has size=%d", len(list))
	}
}

func TestFullTeamReturnsAFullList(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddCharacter(&character.Fake{}, utils.NilCoord)
	}

	list := team.GetCharacters()

	if len(list) != MaxPlayersByTeam {
		t.Errorf("List should be full, but has size=%d", len(list))
	}
}

func TestGetCharactersReturnsValidList(t *testing.T) {
	expectedName := "Bob"
	team := New()
	team.AddCharacter(&character.Fake{FakeName: expectedName}, utils.NilCoord)

	list := team.GetCharacters()

	if list[0].Name() != expectedName {
		t.Errorf("Expected=%v, got=%v", expectedName, list[0].Name())
	}
}

func TestCannotAddCharacterToFullTeam(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddCharacter(&character.Fake{}, utils.NilCoord)
	}

	result := team.AddCharacter(&character.Fake{}, utils.NilCoord)
	expected := false

	if result != expected {
		t.Error("Should not be able to add character to a full team")
	}

	if team.CharactersCount() != MaxPlayersByTeam {
		t.Errorf("team should have %d characters, but has %d", MaxPlayersByTeam, team.CharactersCount())
	}
}

func TestGetAbsentCharacterReturnsNil(t *testing.T) {
	team := New()
	char := &character.Fake{}

	resChar, resPos := team.GetCharacter(char)

	if resChar != nil {
		t.Error("Character should be nil")
	}

	if resPos != utils.NilCoord {
		t.Errorf("Position: expected=%v, got=%v", utils.NilCoord, resPos)
	}
}

func TestIsCharacterAtPosition(t *testing.T) {
	team := New()
	team.AddCharacter(&character.Fake{}, utils.Coord{2, 2})
	team.AddCharacter(&character.Fake{}, utils.Coord{1, 0})

	var tests = []struct {
		c utils.Coord
		b bool
	}{
		{utils.Coord{0, 0}, false},
		{utils.Coord{0, 1}, false},
		{utils.Coord{0, 2}, false},
		{utils.Coord{1, 0}, true},
		{utils.Coord{1, 1}, false},
		{utils.Coord{1, 2}, false},
		{utils.Coord{2, 0}, false},
		{utils.Coord{2, 1}, false},
		{utils.Coord{2, 2}, true},
	}

	for _, test := range tests {
		if team.IsCharacterAtPosition(test.c) != test.b {
			t.Errorf("Expected=%v, result=%v", test.b, team.IsCharacterAtPosition(test.c))
		}
	}
}

func TestCharacterHasCorrectPosition(t *testing.T) {
	team := New()
	c1 := &character.Fake{}
	c2 := &character.Fake{}
	p1 := utils.Coord{2, 3}
	p2 := utils.Coord{1, 4}
	team.AddCharacter(c1, p1)
	team.AddCharacter(c2, p2)

	var tests = []struct {
		char character.Character
		pos  utils.Coord
	}{
		{c1, p1},
		{c2, p2},
	}

	for _, test := range tests {
		char, pos := team.GetCharacter(test.char)
		if char != test.char {
			t.Errorf("Expected=%v, got=%v", test.char, char)
		}
		if pos != test.pos {
			t.Errorf("Expected=%v, got=%v", test.pos, pos)
		}
	}
}

func TestMovedCharacterHasCorrectPosition(t *testing.T) {
	team := New()
	c1 := &character.Fake{}
	c2 := &character.Fake{}
	p1 := utils.Coord{2, 3}
	p2 := utils.Coord{1, 4}
	team.AddCharacter(c1, utils.Coord{0, 0})
	team.AddCharacter(c2, utils.Coord{1, 1})

	team.MoveCharacter(c1, p1)
	team.MoveCharacter(c2, p2)

	var tests = []struct {
		char character.Character
		pos  utils.Coord
	}{
		{c1, p1},
		{c2, p2},
	}

	for _, test := range tests {
		char, pos := team.GetCharacter(test.char)
		if char != test.char {
			t.Errorf("Expected=%v, got=%v", test.char, char)
		}
		if pos != test.pos {
			t.Errorf("Expected=%v, got=%v", test.pos, pos)
		}
	}
}
