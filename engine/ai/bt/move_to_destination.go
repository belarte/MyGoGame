package bt

import (
	"github.com/belarte/MyGoGame/engine/action"
	"github.com/belarte/MyGoGame/engine/ai/pathfinder"
)

// MoveToDestination moves the agent to the position set in the context.
type MoveToDestination struct {
	context    *context
	moveAction action.Action
}

// NewMoveToDestination computes the path to context.destination and returns
// the new task that will perform the MoveAction.
func NewMoveToDestination(context *context) *MoveToDestination {
	return &MoveToDestination{context: context, moveAction: nil}
}

// CheckConditions checks if the MoveAction is doable.
func (task *MoveToDestination) CheckConditions() bool {
	finder := pathfinder.New(task.context.lvl)
	path := finder.ShortestPath(task.context.agent.Position(), task.context.destination)
	task.moveAction = action.NewMoveAction(task.context.queue, task.context.lvl, task.context.agent, &path)
	return task.moveAction.IsDoable()
}

// Perform performs the MoveAction
func (task *MoveToDestination) Perform() Status {
	if task.moveAction.Perform() {
		return success
	}
	return failure
}
