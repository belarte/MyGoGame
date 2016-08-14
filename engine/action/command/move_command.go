package command

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

// MoveCommand mocves an actor at a given position.
type MoveCommand struct {
	commandStatus
	agent       *character.Actor
	destination utils.Coord
	oldPosition utils.Coord
}

// NewMoveCommand return a pointer to a MoveCommand.
func NewMoveCommand(c *character.Actor, pos utils.Coord) *MoveCommand {
	return &MoveCommand{
		commandStatus: newCommandStatus(),
		agent:         c,
		destination:   pos,
		oldPosition:   c.Position(),
	}
}

// Execute moves the actor to the new position.
func (c *MoveCommand) Execute() {
	c.commandStatus.executeIf(func() {
		c.agent.MoveTo(c.destination)
	})
}

// Revert moves the actor to the old position.
func (c *MoveCommand) Revert() {
	c.commandStatus.revertIf(func() {
		c.agent.MoveTo(c.oldPosition)
	})
}
