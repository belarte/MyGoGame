package action

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/action/command"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	. "github.com/belarte/MyGoGame/engine/utils"
)

func TestIsDoableEmptyPath(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Actor{
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	action := NewMoveAction(nil, lvl, char, &level.Path{})

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableNilPath(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Actor{
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	action := NewMoveAction(nil, lvl, char, nil)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoablePathDoesNotStartNextToAgent(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Actor{
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(Coord{0, 4}, 1)
	action := NewMoveAction(nil, lvl, char, &path)

	if action.IsDoable() {
		t.Error("Move action should not be doable")
	}
}

func TestIsDoableOK(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	char := &character.Actor{
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	var path level.Path
	path.Add(Coord{0, 1}, 1)
	action := NewMoveAction(nil, lvl, char, &path)

	if !action.IsDoable() {
		t.Error("Move action should be doable")
	}
}

func TestPerformOk(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	pathCoords := []Coord{Coord{0, 1}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4}}
	char := &character.Actor{
		MovePointsComponent: &character.FakeMovePointsComponent{FakeMovePoints: 10, FakeConsumeMP: true},
		PositionComponent:   &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	queue := &command.FakeQueue{}
	var path level.Path
	for _, coord := range pathCoords {
		path.Add(coord, 1)
	}
	action := NewMoveAction(queue, lvl, char, &path)

	if !action.Perform() {
		t.Error("Move action should have performed.")
	}

	expectedSize := 4
	if queue.Size() != expectedSize {
		t.Errorf("Wrong number of Command. Expected: %d, but is: %d", expectedSize, queue.Size())
	}
}

func TestPerformNotEnoughMovePoints(t *testing.T) {
	lvl := level.New(Coord{1, 5}, 1)
	pathCoords := []Coord{Coord{0, 1}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4}}
	char := &character.Actor{
		MovePointsComponent: &character.FakeMovePointsComponent{FakeMovePoints: 2, FakeConsumeMP: true},
		PositionComponent:   &character.FakePositionComponent{},
	}
	lvl.AddActor(char, Coord{0, 0}, 0)

	queue := &command.FakeQueue{}
	var path level.Path
	for _, coord := range pathCoords {
		path.Add(coord, 1)
	}
	action := NewMoveAction(queue, lvl, char, &path)

	if action.Perform() {
		t.Error("Move action should not have performed successfully.")
	}

	expectedSize := 2
	if queue.Size() != expectedSize {
		t.Errorf("Wrong number of Command. Expected: %d, but is: %d", expectedSize, queue.Size())
	}
}
