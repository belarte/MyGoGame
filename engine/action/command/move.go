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
	cost        float64
}

// NewMove return a pointer to a MoveCommand.
func NewMove(c *character.Actor, pos utils.Coord, cost float64) *Move {
	return &Move{
		status:      newStatus(),
		agent:       c,
		destination: pos,
		oldPosition: c.Position(),
		cost:        cost,
	}
}

// Execute moves the actor to the new position.
func (c *Move) Execute() {
	c.status.executeIf(func() {
		c.agent.MoveTo(c.destination)
		c.agent.ConsumeMovePoints(c.cost)
	})
}

// Revert moves the actor to the old position.
func (c *Move) Revert() {
	c.status.revertIf(func() {
		c.agent.MoveTo(c.oldPosition)
		c.agent.ConsumeMovePoints(-c.cost)
	})
}
