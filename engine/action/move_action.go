package action

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	l "github.com/belarte/MyGoGame/engine/log"
	"github.com/belarte/MyGoGame/engine/utils"
)

// MoveAction move a character along the given path.
type MoveAction struct {
	actionBaseParameters
	path *level.Path
}

// NewMoveAction initialise a new Action.
func NewMoveAction(lvl *level.Level, agent character.Character, path *level.Path) *MoveAction {
	return &MoveAction{newActionBaseParameters(lvl, agent), path}
}

// IsDoable checks if the path is valid.
func (action *MoveAction) IsDoable() bool {
	if action.path == nil || action.path.Size() == 0 {
		l.Log(action.agent.Name() + " cannot move: empty path.")
		return false
	}

	if !utils.AreAdjacent(action.level.PositionOf(action.agent), action.path.Path[0].Coord) {
		l.Log("Inconsistant path: does not start next to agent.")
		return false
	}

	l.Log(action.agent.Name() + " can move.")
	return true
}

// Perform move the agent along the path.
// Return true if the agent reached destination.
// Return false if the agent lacked MP to reach destination
// or if an event occured while moving.
func (action *MoveAction) Perform() bool {
	for _, step := range action.path.Path {
		consumed := action.agent.ConsumeMovePoints(step.Cost)
		if !consumed {
			l.Log(action.agent.Name() + "  does not have enough move points, action terminated. ")
			break
		}

		action.team.MoveCharacter(action.agent, step.Coord)
		// TODO: implement events
	}

	result := action.level.PositionOf(action.agent) == action.path.Path[len(action.path.Path)-1].Coord
	if result {
		l.Log(action.agent.Name() + " arrived at destination.")
	} else {
		l.Log("Something happened to " + action.agent.Name() + " while moving.")
	}
	return result
}
