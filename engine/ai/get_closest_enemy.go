package ai

import (
	. "github.com/belarte/MyGoGame/engine/utils"
	"math"
)

type GetClosestEnemy struct {
	context *Context
}

func NewGetClosestEnemies(context *Context) *GetClosestEnemy {
	context.closestEnemyPosition = NilCoord
	return &GetClosestEnemy{context}
}

func (self *GetClosestEnemy) CheckConditions() bool {
	return len(self.context.visibleEnemies) > 0
}

func (self *GetClosestEnemy) Perform() bool {
	distance := math.MaxFloat64
	for _, opponent := range self.context.visibleEnemies {
		if opponent.dist < distance {
			self.context.closestEnemy = opponent.char
			self.context.closestEnemyPosition = opponent.pos
			distance = opponent.dist
		}
	}

	return true
}
