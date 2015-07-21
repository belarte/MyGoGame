package ai

import (
	. "github.com/belarte/MyGoGame/engine/utils"
)

type GetVisibleEnemies struct {
	context *context
}

func NewGetVisibleEnemies(context *context) *GetVisibleEnemies {
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
