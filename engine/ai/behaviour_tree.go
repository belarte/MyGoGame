package ai

import (
	"github.com/belarte/MyGoGame/engine/core"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

type charPosDist struct {
	char character.Character
	pos  utils.Coord
	dist float64
}

type context struct {
	level                *core.Level
	agent                character.Character
	positionOfAgent      utils.Coord
	visibleEnemies       []charPosDist
	closestEnemy         character.Character
	closestEnemyPosition utils.Coord
	destination          utils.Coord
}

func newContext(level *core.Level, agent character.Character) *context {
	return &context{
		level:                level,
		agent:                agent,
		positionOfAgent:      level.PositionOf(agent),
		visibleEnemies:       make([]charPosDist, 0, 4),
		closestEnemyPosition: utils.NilCoord,
		destination:          utils.NilCoord,
	}
}

// Task represents an abstract task.
// A task is composed by two methods:
// CheckConditions and Perform.
type Task interface {
	CheckConditions() bool
	Perform() bool
}

// MockTask is a mock implementation of a task.
type MockTask struct {
	CheckConditionsMock bool
	PerformMock         bool
}

// CheckConditions implementation for MockTask
func (mock *MockTask) CheckConditions() bool {
	return mock.CheckConditionsMock
}

// Perform implementation for MockTask.
func (mock *MockTask) Perform() bool {
	return mock.PerformMock
}
