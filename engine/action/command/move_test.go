package command

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestExecuteDoesMoveAgentToDestination(t *testing.T) {
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent: &character.Position2DComponent{},
	}

	command := NewMove(c, to)
	command.Execute()

	if c.Position() != to {
		t.Errorf("Expected position is %+v, but actually is %+v", to, c.Position())
	}
}

func TestRevertDoesNothingIfCommandHasNotBeenExecuted(t *testing.T) {
	c := &character.Actor{
		PositionComponent: &character.Position2DComponent{},
	}
	from := c.Position()
	to := utils.Coord{X: 1, Y: 1}

	command := NewMove(c, to)
	command.Revert()

	if c.Position() != from {
		t.Errorf("Expected position is %+v, but actually is %+v", from, c.Position())
	}
}

func TestRevertDoesMoveAgentToInitialPosition(t *testing.T) {
	from := utils.Coord{X: 2, Y: 2}
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent: &character.Position2DComponent{},
	}
	c.MoveTo(from)

	command := NewMove(c, to)
	command.Execute()
	command.Revert()

	if c.Position() != from {
		t.Errorf("Expected position is %+v, but actually is %+v", from, c.Position())
	}
}

func TestExecuteTwiceDoesNothing(t *testing.T) {
	to := utils.Coord{X: 1, Y: 1}
	c := &character.Actor{
		PositionComponent: &character.Position2DComponent{},
	}

	command := NewMove(c, to)
	command.Execute()
	command.Execute()

	if c.Position() != to {
		t.Errorf("Expected position is %+v, but actually is %+v", to, c.Position())
	}
}
