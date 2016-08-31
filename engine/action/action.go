// Package action implements actions doable by a Actor.
package action

import (
	"github.com/belarte/MyGoGame/engine/action/command"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/core/team"
)

type actionBaseParameters struct {
	queue command.Queue
	level *level.Level
	agent *character.Actor
	team  *team.Team
}

func newActionBaseParameters(queue command.Queue, level *level.Level, agent *character.Actor) actionBaseParameters {
	team := level.GetTeamOf(agent)
	return actionBaseParameters{queue, level, agent, team}
}

// Action defines an action that is doable by a character on a level.
type Action interface {
	IsDoable() bool
	Perform() bool
}

// Fake action
type Fake struct {
	FakeIsDoable, FakePerform bool
}

// IsDoable return the parameter given on initialisation.
func (action *Fake) IsDoable() bool {
	return action.FakeIsDoable
}

// Perform return the parameter given on initialisation.
func (action *Fake) Perform() bool {
	return action.FakePerform
}
