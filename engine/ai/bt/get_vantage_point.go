package bt

import (
	"math"

	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

// GetVantagePoint computes the safest place where an attack is possible
// on the closest enemy. The computed position will be at range and the
// enemy will be reachable from an attack in line.
type GetVantagePoint struct {
	context *context
}

// NewGetVantagePoint returns the new task.
func NewGetVantagePoint(context *context) *GetVantagePoint {
	return &GetVantagePoint{context}
}

// CheckConditions checks that the level, the agent and the closest
// enemy's position are not nil.
func (task *GetVantagePoint) CheckConditions() bool {
	return task.context.lvl != nil &&
		task.context.agent != nil &&
		task.context.closestEnemyPosition != utils.NilCoord
}

// Perform computes the targeted position. It Computes the circle with radius agent.range
// from where the target will be reachable for a direct attack.
func (task *GetVantagePoint) Perform() bool {
	positions := utils.Circle(task.context.closestEnemyPosition, task.context.agent.Range())

	distance := math.MaxFloat64
	task.context.destination = utils.NilCoord
	maps := task.context.lvl.Map()
	for _, position := range positions {
		if !maps.IsWithinBounds(position) || task.isSightBlocked(maps, position) {
			continue
		}

		distToAgent := utils.Distance(task.context.positionOfAgent, position)
		if distToAgent < distance {
			distance = distToAgent
			task.context.destination = position
		}
	}

	return task.context.destination != utils.NilCoord
}

func (task *GetVantagePoint) isSightBlocked(maps *level.Map, pos utils.Coord) bool {
	line := utils.Line(task.context.closestEnemyPosition, pos)
	for _, pos := range line {
		if maps.GetCell(pos) == level.WALL {
			return true
		}
	}

	return false
}
