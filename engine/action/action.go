package action

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
)

type ActionBaseParameters struct {
	level *Level
	agent *Character
	team  *Team
	logs  [2]string
}

func NewActionBaseParameters(level *Level, agent *Character) ActionBaseParameters {
	team := level.GetTeamOf(agent)
	return ActionBaseParameters{level, agent, team, [2]string{"", ""}}
}

type Action interface {
	IsDoable() bool
	Perform() bool
}

type MoveAction struct {
	ActionBaseParameters
	path *Path
}

func NewMoveAction(lvl *Level, agent *Character, path *Path) *MoveAction {
	return &MoveAction{NewActionBaseParameters(lvl, agent), path}
}

func (self *MoveAction) IsDoable() bool {
	if self.path == nil || self.path.Size() == 0 {
		self.logs[0] = self.agent.Name() + " cannot move: empty path"
		return false
	}

	self.logs[0] = self.agent.Name() + " can move."
	return true
}

func (self *MoveAction) Perform() bool {
	movePoints := float64(self.agent.MovePoints())
	for _, step := range self.path.Path {
		if movePoints < self.path.Cost() {
			self.logs[1] += self.agent.Name() + "  does not have enough move points, action terminated. "
			break
		}

		self.team.MoveCharacter(self.agent, step.Coord)
		movePoints -= self.path.Cost()
		// TODO: implement events
	}

	result := EqualCoord(self.level.PositionOf(self.agent), self.path.Path[len(self.path.Path)-1].Coord)
	if result {
		self.logs[1] += self.agent.Name() + " arrived at destination."
	} else {
		self.logs[1] += "Something happened to " + self.agent.Name() + " while moving."
	}
	return result
}

type AttackAction struct {
	ActionBaseParameters
	target *Character
}

func NewAttackAction(lvl *Level, agent, target *Character) *AttackAction {
	return &AttackAction{NewActionBaseParameters(lvl, agent), target}
}

func (self *AttackAction) IsDoable() bool {
	return false
}

func (self *AttackAction) Perform() bool {
	return false
}
