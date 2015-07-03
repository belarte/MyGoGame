package action

import (
	"github.com/belarte/MyGoGame/engine/core"
)

type ActionBaseParameters struct {
	level *core.Level
	agent *core.Character
	team  *core.Team
	logs  [2]string
}

func NewActionBaseParameters(level *core.Level, agent *core.Character) ActionBaseParameters {
	team := level.GetTeamOfCharacter(agent)
	return ActionBaseParameters{level, agent, team, [2]string{"", ""}}
}

type Action interface {
	IsDoable() bool
	Perform()
}

type MoveAction struct {
	ActionBaseParameters
	path core.Path
}

func NewMoveAction(lvl *core.Level, agent *core.Character, path core.Path) *MoveAction {
	return &MoveAction{NewActionBaseParameters(lvl, agent), path}
}

func (self *MoveAction) IsDoable() bool {
	if self.path.Size() == 0 {
		self.logs[0] = self.agent.Name() + " cannot move: empty path"
		return false
	}

	if self.agent.MovePoints() < int(self.path.Cost()) {
		self.logs[0] = self.agent.Name() + " cannot move: not enough move points"
		return false
	}

	self.logs[0] = self.agent.Name() + " can move."
	return true
}

func (self *MoveAction) Perform() {
	for _, step := range self.path.Path {
		self.team.MoveCharacter(self.agent, step.Coord)
		// TODO: implement cost
		// TODO: implement events
		// TODO: log
	}

	self.logs[1] = self.agent.Name() + " arrived at destination."
}

type AttackAction struct {
	ActionBaseParameters
	target *core.Character
}

func NewAttackAction(lvl *core.Level, agent, target *core.Character) *AttackAction {
	return &AttackAction{NewActionBaseParameters(lvl, agent), target}
}

func (self *AttackAction) IsDoable() bool {
	return false
}

func (self *AttackAction) Perform() {
}
