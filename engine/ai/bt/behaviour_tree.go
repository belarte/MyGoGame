package bt

import (
	"github.com/belarte/MyGoGame/engine/action/command"
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

type charDist struct {
	char *character.Actor
	dist float64
}

type context struct {
	queue                command.Queue
	lvl                  *level.Level
	agent                *character.Actor
	visibleEnemies       []charDist
	closestEnemy         *character.Actor
	closestEnemyPosition utils.Coord
	destination          utils.Coord
}

func newContext(lvl *level.Level, agent *character.Actor) *context {
	return &context{
		lvl:                  lvl,
		agent:                agent,
		visibleEnemies:       make([]charDist, 0, 4),
		closestEnemyPosition: utils.NilCoord,
		destination:          utils.NilCoord,
	}
}

type Status int

const (
	ready Status = iota
	running
	success
	failure
)

// Task represents an abstract task.
// A task is composed by two methods:
// CheckConditions and Perform.
type Task interface {
	CheckConditions() bool
	Perform() Status
}

// Fake is a task implementation for testing purposes.
type Fake struct {
	FakeCheckConditions bool
	FakePerform         Status
}

// CheckConditions implementation for MockTask
func (task *Fake) CheckConditions() bool {
	return task.FakeCheckConditions
}

// Perform implementation for MockTask.
func (task *Fake) Perform() Status {
	return task.FakePerform
}
