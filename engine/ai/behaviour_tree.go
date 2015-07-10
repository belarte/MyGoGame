package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
)

type charPosDist struct {
	char *Character
	pos  *Coord
	dist float64
}

type Context struct {
	level                *Level
	agent                *Character
	positionOfAgent      Coord
	visibleEnemies       []charPosDist
	closestEnemy         *Character
	closestEnemyPosition *Coord
}

func NewContext(level *Level, agent *Character) *Context {
	return &Context{level, agent, level.PositionOf(agent), make([]charPosDist, 0, 4), nil, nil}
}

type Task interface {
	CheckConditions()
	Perform() bool
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
	self.context.visibleEnemies = make([]charPosDist, 0, 4)

	if len(opponents) == 0 {
		return false
	}

	for _, opponent := range opponents {
		position := self.context.level.PositionOf(opponent)
		distance := Distance(self.context.positionOfAgent, position)
		if distance <= float64(self.context.agent.Visibility()) {
			char := charPosDist{opponent, &position, distance}
			self.context.visibleEnemies = append(self.context.visibleEnemies, char)
		}
	}

	return len(self.context.visibleEnemies) > 0
}

type GetClosestEnemy struct {
	context *Context
}

func NewGetClosestEnemies(context *Context) *GetClosestEnemy {
	return &GetClosestEnemy{context}
}

func (self *GetClosestEnemy) CheckConditions() bool {
	return len(self.context.visibleEnemies) > 0
}

func (self *GetClosestEnemy) Perform() bool {
	distance := 123456789.0
	for _, opponent := range self.context.visibleEnemies {
		if opponent.dist < distance {
			self.context.closestEnemy = opponent.char
			self.context.closestEnemyPosition = opponent.pos
			distance = opponent.dist
		}
	}

	return true
}
