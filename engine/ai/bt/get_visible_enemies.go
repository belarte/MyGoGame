package bt

import (
	"github.com/belarte/MyGoGame/engine/utils"
)

// GetVisibleEnemies gets all enemies that are visible by the agent
// or one of its team member.
type GetVisibleEnemies struct {
	context *context
}

// NewGetVisibleEnemies returns the new task.
func NewGetVisibleEnemies(context *context) *GetVisibleEnemies {
	return &GetVisibleEnemies{context: context}
}

// CheckConditions checks that the agent and the level are not nil.
func (task *GetVisibleEnemies) CheckConditions() bool {
	return task.context.lvl != nil && task.context.agent != nil
}

// Perform gets the visible enemies.
// It fails if no visible enemies are found.
func (task *GetVisibleEnemies) Perform() Status {
	opponents := task.context.lvl.GetOpponentsOf(task.context.agent)
	task.context.visibleEnemies = make([]charDist, 0, 4)

	if len(opponents) == 0 {
		return failure
	}

	for _, opponent := range opponents {
		position := opponent.Position()
		distance := utils.Distance(task.context.agent.Position(), position)
		if distance <= float64(task.context.agent.Visibility()) {
			if task.isEnemyAtPositionVisible(position) {
				char := charDist{opponent, distance}
				task.context.visibleEnemies = append(task.context.visibleEnemies, char)
			}
		}
	}

	if len(task.context.visibleEnemies) > 0 {
		return success
	}
	return failure
}

func (task *GetVisibleEnemies) isEnemyAtPositionVisible(pos utils.Coord) bool {
	path := utils.Line(task.context.agent.Position(), pos)

	for _, coord := range path {
		if task.context.lvl.IsObstacleAtPosition(coord) {
			return false
		}
	}

	return true
}
