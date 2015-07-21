package ai

import (
	. "github.com/belarte/MyGoGame/engine/action"
)

type MoveToDestination struct {
	context    *context
	moveAction Action
}

func NewMoveToDestination(context *context) *MoveToDestination {
	finder := NewPathFinder(context.level)
	path := finder.ShortestPath(context.positionOfAgent, context.destination)
	action := NewMoveAction(context.level, context.agent, &path)
	return &MoveToDestination{context: context, moveAction: action}
}

func (self *MoveToDestination) CheckConditions() bool {
	return self.moveAction.IsDoable()
}

func (self *MoveToDestination) Perform() bool {
	return self.moveAction.Perform()
}
