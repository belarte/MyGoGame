package command

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Move mocves an actor at a given position.
type Move struct {
	status
	agent       *character.Actor
	destination utils.Coord
	oldPosition utils.Coord
}

// NewMove return a pointer to a MoveCommand.
func NewMove(c *character.Actor, pos utils.Coord) *Move {
	return &Move{
		status:      newStatus(),
		agent:       c,
		destination: pos,
		oldPosition: c.Position(),
	}
}

// Execute moves the actor to the new position.
func (c *Move) Execute() {
	c.status.executeIf(func() {
		c.agent.MoveTo(c.destination)
	})
}

// Revert moves the actor to the old position.
func (c *Move) Revert() {
	c.status.revertIf(func() {
		c.agent.MoveTo(c.oldPosition)
	})
}
