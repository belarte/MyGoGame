// Package action implements actions doable by a Character.
package action

import (
	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
)

type actionBaseParameters struct {
	level *core.Level
	agent character.Character
	team  *core.Team
	logs  [2]string
}

func newactionBaseParameters(level *core.Level, agent character.Character) actionBaseParameters {
	team := level.GetTeamOf(agent)
	return actionBaseParameters{level, agent, team, [2]string{"", ""}}
}

// Action defines an action that is doable by a character on a level.
type Action interface {
	IsDoable() bool
	Perform() bool
}

// Mock action
type Mock struct {
	IsDoableMock, PerformMock bool
}

// IsDoable return the parameter given on initialisation.
func (action *Mock) IsDoable() bool {
	return action.IsDoableMock
}

// Perform return the parameter given on initialisation.
func (action *Mock) Perform() bool {
	return action.PerformMock
}

/*
type AttackAction struct {
	actionBaseParameters
	target core.Character
}

func NewAttackAction(lvl *core.Level, agent, target core.Character) *AttackAction {
	return &AttackAction{newactionBaseParameters(lvl, agent), target}
}

func (action *AttackAction) IsDoable() bool {
	return false
}

func (action *AttackAction) Perform() bool {
	return false
}
*/
