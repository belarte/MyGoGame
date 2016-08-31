package action

import (
	"github.com/belarte/MyGoGame/engine/action/command"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

// MoveAction move a character along the given path.
type MoveAction struct {
	actionBaseParameters
	path *level.Path
}

// NewMoveAction initialise a new Action.
func NewMoveAction(q command.Queue, lvl *level.Level, agent *character.Actor, path *level.Path) *MoveAction {
	return &MoveAction{newActionBaseParameters(q, lvl, agent), path}
}

// IsDoable checks if the path is valid.
func (action *MoveAction) IsDoable() bool {
	if action.path == nil || action.path.Size() == 0 {
		return false
	}

	if !utils.AreAdjacent(action.agent.Position(), action.path.Path[0].Coord) {
		return false
	}

	return true
}

// Perform move the agent along the path.
// Return true if the agent reached destination.
// Return false if the agent lacked MP to reach destination
// or if an event occured while moving.
func (action *MoveAction) Perform() bool {
	movePoints := action.agent.MovePoints()
	aggregatedMovePoints := 0.
	for _, step := range action.path.Path {
		aggregatedMovePoints += step.Cost
		if aggregatedMovePoints > movePoints {
			break
		}

		c := command.NewMove(action.agent, step.Coord, step.Cost)
		action.queue.Add(c)
	}

	return aggregatedMovePoints < movePoints
}
