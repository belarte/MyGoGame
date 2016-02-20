// Package action implements actions doable by a Character.
package action

import (
	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/team"
)

type actionBaseParameters struct {
	level *core.Level
	agent character.Character
	team  *team.Team
}

func newActionBaseParameters(level *core.Level, agent character.Character) actionBaseParameters {
	team := level.GetTeamOf(agent)
	return actionBaseParameters{level, agent, team}
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
