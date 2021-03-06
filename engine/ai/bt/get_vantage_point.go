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
func (task *GetVantagePoint) Perform() Status {
	positions := utils.Circle(task.context.closestEnemyPosition, task.context.agent.Range())

	distance := math.MaxFloat64
	task.context.destination = utils.NilCoord
	lvl := task.context.lvl
	for _, position := range positions {
		if !lvl.IsWithinBounds(position) || task.isSightBlocked(lvl, position) {
			continue
		}

		distToAgent := utils.Distance(task.context.agent.Position(), position)
		if distToAgent < distance {
			distance = distToAgent
			task.context.destination = position
		}
	}

	if task.context.destination == utils.NilCoord {
		return failure
	}
	return success
}

func (task *GetVantagePoint) isSightBlocked(lvl *level.Level, pos utils.Coord) bool {
	line := utils.Line(task.context.closestEnemyPosition, pos)
	for _, pos := range line {
		if lvl.GetCell(pos) == level.WallCell {
			return true
		}
	}

	return false
}
