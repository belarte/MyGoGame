package ai

import (
	"github.com/belarte/MyGoGame/engine/action"
)

// MoveToDestination moves the agent to the position set in the context.
type MoveToDestination struct {
	context    *context
	moveAction action.Action
}

// NewMoveToDestination computes the path to context.destination and returns
// the new task that will perform the MoveAction.
func NewMoveToDestination(context *context) *MoveToDestination {
	finder := NewPathFinder(context.level)
	path := finder.ShortestPath(context.positionOfAgent, context.destination)
	action := action.NewMoveAction(context.level, context.agent, &path)
	return &MoveToDestination{context: context, moveAction: action}
}

// CheckConditions checks if the MoveAction is doable.
func (task *MoveToDestination) CheckConditions() bool {
	return task.moveAction.IsDoable()
}

// Perform performs the MoveAction
func (task *MoveToDestination) Perform() bool {
	return task.moveAction.Perform()
}
