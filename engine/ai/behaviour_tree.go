package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
)

type Context struct {
	level           *Level
	agent           *Character
	positionOfAgent Coord
	visibleEnemies  []*Character
}

func NewContext(level *Level, agent *Character) *Context {
	return &Context{level, agent, level.PositionOf(agent), make([]*Character, 0, 4)}
}

type GetVisibleEnemies struct {
	context *Context
}

func NewGetVisibleEnemies(context *Context) *GetVisibleEnemies {
	return &GetVisibleEnemies{context}
}

func (self *GetVisibleEnemies) CheckConditions() bool {
	return self.context.level != nil && self.context.agent != nil
}

func (self *GetVisibleEnemies) Perform() bool {
	opponents := self.context.level.GetOpponentsOf(self.context.agent)
	self.context.visibleEnemies = make([]*Character, 0, 4)

	if len(opponents) == 0 {
		return false
	}

	for _, opponent := range opponents {
		position := self.context.level.PositionOf(opponent)
		distance := Distance(self.context.positionOfAgent, position)
		if distance <= float64(self.context.agent.Visibility()) {
			self.context.visibleEnemies = append(self.context.visibleEnemies, opponent)
		}
	}

	return len(self.context.visibleEnemies) > 0
}
