package command

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestExecuteDoesMoveAgentToDestination(t *testing.T) {
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent:   &character.Position2DComponent{},
		MovePointsComponent: &character.FakeMovePointsComponent{},
	}

	command := NewMove(c, to, 0)
	command.Execute()

	if c.Position() != to {
		t.Errorf("Expected position is %+v, but actually is %+v", to, c.Position())
	}
}

func TestRevertDoesNothingIfCommandHasNotBeenExecuted(t *testing.T) {
	c := &character.Actor{
		PositionComponent:   &character.Position2DComponent{},
		MovePointsComponent: &character.FakeMovePointsComponent{},
	}
	from := c.Position()
	to := utils.Coord{X: 1, Y: 1}

	command := NewMove(c, to, 0)
	command.Revert()

	if c.Position() != from {
		t.Errorf("Expected position is %+v, but actually is %+v", from, c.Position())
	}
}

func TestRevertDoesMoveAgentToInitialPosition(t *testing.T) {
	from := utils.Coord{X: 2, Y: 2}
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent:   &character.Position2DComponent{},
		MovePointsComponent: &character.FakeMovePointsComponent{},
	}
	c.MoveTo(from)

	command := NewMove(c, to, 0)
	command.Execute()
	command.Revert()

	if c.Position() != from {
		t.Errorf("Expected position is %+v, but actually is %+v", from, c.Position())
	}
}

func TestExecuteTwiceDoesNothing(t *testing.T) {
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent:   &character.Position2DComponent{},
		MovePointsComponent: &character.FakeMovePointsComponent{},
	}

	command := NewMove(c, to, 0)
	command.Execute()
	command.Execute()

	if c.Position() != to {
		t.Errorf("Expected position is %+v, but actually is %+v", to, c.Position())
	}
}

func TestExecuteConsumeMovePoints(t *testing.T) {
	c := &character.Actor{
		PositionComponent:   &character.FakePositionComponent{},
		MovePointsComponent: character.NewSimpleMovePointsComponent(5, 0),
	}

	command := NewMove(c, utils.NilCoord, 2)
	command.Execute()

	if c.MovePoints() != 3 {
		t.Errorf("Expected move points: 3, got: %f", c.MovePoints())
	}
}

func TestExecuteThenRevertRestoreMovePoints(t *testing.T) {
	c := &character.Actor{
		PositionComponent:   &character.FakePositionComponent{},
		MovePointsComponent: character.NewSimpleMovePointsComponent(5, 0),
	}

	command := NewMove(c, utils.NilCoord, 2)
	command.Execute()
	command.Revert()

	if c.MovePoints() != 5 {
		t.Errorf("Expected move points: 3, got: %f", c.MovePoints())
	}
}
