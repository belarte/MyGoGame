package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
)

type charPosDist struct {
	char *Character
	pos  Coord
	dist float64
}

type Context struct {
	level                *Level
	agent                *Character
	positionOfAgent      Coord
	visibleEnemies       []charPosDist
	closestEnemy         *Character
	closestEnemyPosition Coord
	destination          Coord
}

func NewContext(level *Level, agent *Character) *Context {
	return &Context{level: level, agent: agent, positionOfAgent: level.PositionOf(agent)}
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
			char := charPosDist{opponent, position, distance}
			self.context.visibleEnemies = append(self.context.visibleEnemies, char)
		}
	}

	return len(self.context.visibleEnemies) > 0
}

type GetClosestEnemy struct {
	context *Context
}

func NewGetClosestEnemies(context *Context) *GetClosestEnemy {
	context.closestEnemyPosition = Coord{-1, -1}
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

type GetVantagePoint struct {
	context *Context
}

func NewGetVantagePoint(context *Context) *GetVantagePoint {
	return &GetVantagePoint{context}
}

func (self *GetVantagePoint) CheckConditions() bool {
	return self.context.level != nil &&
		self.context.agent != nil &&
		self.context.closestEnemyPosition != Coord{-1, -1}
}

func (self *GetVantagePoint) Perform() bool {
	positions := Circle(self.context.closestEnemyPosition, self.context.agent.Range())

	distance := 123456789.0
	self.context.destination = Coord{-1, -1}
	maps := self.context.level.Map()
	for _, position := range positions {
		if !maps.IsWithinBounds(position) || self.isSightBlocked(maps, position) {
			continue
		}

		distToAgent := Distance(self.context.positionOfAgent, position)
		if distToAgent < distance {
			distance = distToAgent
			self.context.destination = position
		}
	}

	return self.context.destination != Coord{-1, -1}
}

func (self *GetVantagePoint) isSightBlocked(maps *Map, pos Coord) bool {
	line := Line(self.context.closestEnemyPosition, pos)
	for _, pos := range line {
		if maps.GetCell(pos) == WALL {
			return true
		}
	}

	return false
}
