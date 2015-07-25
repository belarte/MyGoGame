package action

import (
	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

// MoveAction move a character along the given path.
type MoveAction struct {
	actionBaseParameters
	path *core.Path
}

// NewMoveAction initialise a new Action.
func NewMoveAction(lvl *core.Level, agent character.Character, path *core.Path) *MoveAction {
	return &MoveAction{newactionBaseParameters(lvl, agent), path}
}

// IsDoable checks if the path is valid.
func (action *MoveAction) IsDoable() bool {
	if action.path == nil || action.path.Size() == 0 {
		action.logs[0] = action.agent.Name() + " cannot move: empty path."
		return false
	}

	if !utils.AreAdjacent(action.level.PositionOf(action.agent), action.path.Path[0].Coord) {
		action.logs[0] = "Inconsistant path: does not start next to agent."
		return false
	}

	action.logs[0] = action.agent.Name() + " can move."
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
			action.logs[1] += action.agent.Name() + "  does not have enough move points, action terminated. "
			break
		}

		action.team.MoveCharacter(action.agent, step.Coord)
		// TODO: implement events
	}

	result := action.level.PositionOf(action.agent) == action.path.Path[len(action.path.Path)-1].Coord
	if result {
		action.logs[1] += action.agent.Name() + " arrived at destination."
	} else {
		action.logs[1] += "Something happened to " + action.agent.Name() + " while moving."
	}
	return result
}
