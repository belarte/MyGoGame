package team

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestEmptyTeamIsEmpty(t *testing.T) {
	team := New()

	size := team.ActorsCount()
	expectedSize := 0

	if size != expectedSize {
		t.Errorf("Expected=%d, Result=%d", expectedSize, size)
	}
}

func TestFullTeamIsFull(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddActor(&character.Actor{
			PositionComponent: &character.FakePositionComponent{FakePosition: utils.NilCoord},
		})
	}

	if !team.IsFull() {
		t.Error("team should be full")
	}

	if team.ActorsCount() != MaxPlayersByTeam {
		t.Errorf("team should have %d characters, but has %d", MaxPlayersByTeam, team.ActorsCount())
	}
}

func TestEmptyTeamReturnsEmptyListOfActors(t *testing.T) {
	team := New()
	list := team.GetActors()

	if len(list) != 0 {
		t.Errorf("List should be empty, but has size=%d", len(list))
	}
}

func TestFullTeamReturnsAFullList(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddActor(&character.Actor{
			PositionComponent: &character.FakePositionComponent{FakePosition: utils.NilCoord},
		})
	}

	list := team.GetActors()

	if len(list) != MaxPlayersByTeam {
		t.Errorf("List should be full, but has size=%d", len(list))
	}
}

func TestGetActorsReturnsValidList(t *testing.T) {
	expectedName := "Bob"
	team := New()
	team.AddActor(&character.Actor{
		StatsComponent:    &character.FakeStatsComponent{FakeName: expectedName},
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.NilCoord},
	})

	list := team.GetActors()

	if list[0].Name() != expectedName {
		t.Errorf("Expected=%v, got=%v", expectedName, list[0].Name())
	}
}

func TestCannotAddActorToFullTeam(t *testing.T) {
	team := New()
	for i := 0; i < MaxPlayersByTeam; i++ {
		team.AddActor(&character.Actor{
			PositionComponent: &character.FakePositionComponent{FakePosition: utils.NilCoord},
		})
	}

	result := team.AddActor(&character.Actor{
		PositionComponent: &character.FakePositionComponent{FakePosition: utils.NilCoord},
	})
	expected := false

	if result != expected {
		t.Error("Should not be able to add character to a full team")
	}

	if team.ActorsCount() != MaxPlayersByTeam {
		t.Errorf("team should have %d characters, but has %d", MaxPlayersByTeam, team.ActorsCount())
	}
}

//TODO test coverage
//func TestIsActorAtPosition(t *testing.T) {
//	team := New()
//	team.AddActor(&character.Fake{
//		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 2, Y: 2}},
//	})
//	team.AddActor(&character.Fake{
//		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.Coord{X: 1, Y: 0}},
//	})
//
//	var tests = []struct {
//		c utils.Coord
//		b bool
//	}{
//		{utils.Coord{0, 0}, false},
//		{utils.Coord{0, 1}, false},
//		{utils.Coord{0, 2}, false},
//		{utils.Coord{1, 0}, true},
//		{utils.Coord{1, 1}, false},
//		{utils.Coord{1, 2}, false},
//		{utils.Coord{2, 0}, false},
//		{utils.Coord{2, 1}, false},
//		{utils.Coord{2, 2}, true},
//	}
//
//	for _, test := range tests {
//		if team.IsActorAtPosition(test.c) != test.b {
//			t.Errorf("Expected=%v, result=%v", test.b, team.IsActorAtPosition(test.c))
//		}
//	}
//}
//
//func TestActorHasCorrectPosition(t *testing.T) {
//	team := New()
//	c1 := &character.Fake{}
//	c2 := &character.Fake{}
//	p1 := utils.Coord{2, 3}
//	p2 := utils.Coord{1, 4}
//	c1.MoveTo(p1)
//	c2.MoveTo(p2)
//	team.AddActor(c1)
//	team.AddActor(c2)
//
//	var tests = []struct {
//		char character.Actor
//		pos  utils.Coord
//	}{
//		{c1, p1},
//		{c2, p2},
//	}
//
//	for _, test := range tests {
//		char, pos := team.GetActor(test.char)
//		if char != test.char {
//			t.Errorf("Expected=%v, got=%v", test.char, char)
//		}
//		if pos != test.pos {
//			t.Errorf("Expected=%v, got=%v", test.pos, pos)
//		}
//	}
//}
//
//func TestMovedActorHasCorrectPosition(t *testing.T) {
//	team := New()
//	c1 := &character.Fake{
//		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.NilCoord},
//	}
//	c2 := &character.Fake{
//		FakePositionComponent: character.FakePositionComponent{FakePosition: utils.NilCoord},
//	}
//	p1 := utils.Coord{2, 3}
//	p2 := utils.Coord{1, 4}
//	team.AddActor(c1)
//	team.AddActor(c2)
//
//	team.MoveActor(c1, p1)
//	team.MoveActor(c2, p2)
//
//	var tests = []struct {
//		char character.Actor
//		pos  utils.Coord
//	}{
//		{c1, p1},
//		{c2, p2},
//	}
//
//	for _, test := range tests {
//		char, pos := team.GetActor(test.char)
//		if char != test.char {
//			t.Errorf("Expected=%v, got=%v", test.char, char)
//		}
//		if pos != test.pos {
//			t.Errorf("Expected=%v, got=%v", test.pos, pos)
//		}
//	}
//}
