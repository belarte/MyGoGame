package bt

import (
	"math"

	"github.com/belarte/MyGoGame/engine/utils"
)

// GetClosestEnemy computes the closest enemy.
type GetClosestEnemy struct {
	context *context
}

// NewGetClosestEnemies returns the new task.
func NewGetClosestEnemies(context *context) *GetClosestEnemy {
	context.closestEnemyPosition = utils.NilCoord
	return &GetClosestEnemy{context}
}

// CheckConditions checks thqt a least one enemy is visible.
func (task *GetClosestEnemy) CheckConditions() bool {
	return len(task.context.visibleEnemies) > 0
}

// Perform finds the enemy with the shortest Euclidian distance
// from the agent.
func (task *GetClosestEnemy) Perform() Status {
	distance := math.MaxFloat64
	for _, opponent := range task.context.visibleEnemies {
		if opponent.dist < distance {
			task.context.closestEnemy = opponent.char
			task.context.closestEnemyPosition = opponent.char.Position()
			distance = opponent.dist
		}
	}

	return success
}
