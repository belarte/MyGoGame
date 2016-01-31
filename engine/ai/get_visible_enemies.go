package ai

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
	return task.context.level != nil && task.context.agent != nil
}

// Perform gets the visible enemies.
// It fails if no visible enemies are found.
func (task *GetVisibleEnemies) Perform() bool {
	opponents := task.context.level.GetOpponentsOf(task.context.agent)
	task.context.visibleEnemies = make([]charPosDist, 0, 4)

	if len(opponents) == 0 {
		return false
	}

	for _, opponent := range opponents {
		position := task.context.level.PositionOf(opponent)
		distance := utils.Distance(task.context.positionOfAgent, position)
		if distance <= float64(task.context.agent.Visibility()) {
			if task.isEnemyAtPositionVisible(position) {
				char := charPosDist{opponent, position, distance}
				task.context.visibleEnemies = append(task.context.visibleEnemies, char)
			}
		}
	}

	return len(task.context.visibleEnemies) > 0
}

func (task *GetVisibleEnemies) isEnemyAtPositionVisible(pos utils.Coord) bool {
	path := utils.Line(task.context.positionOfAgent, pos)

	for _, coord := range path {
		if task.context.level.IsObstacleAtPosition(coord) {
			return false
		}
	}

	return true
}
