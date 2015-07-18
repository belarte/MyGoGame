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
	return &Context{
		level:                level,
		agent:                agent,
		positionOfAgent:      level.PositionOf(agent),
		closestEnemyPosition: NilCoord,
		destination:          NilCoord,
	}
}

type Task interface {
	CheckConditions() bool
	Perform() bool
}

type MockTask struct {
	CheckConditionsMock bool
	PerformMock         bool
}

func (self *MockTask) CheckConditions() bool {
	return self.CheckConditionsMock
}

func (self *MockTask) Perform() bool {
	return self.PerformMock
}
