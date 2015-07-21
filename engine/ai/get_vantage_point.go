package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"math"
)

type GetVantagePoint struct {
	context *context
}

func NewGetVantagePoint(context *context) *GetVantagePoint {
	return &GetVantagePoint{context}
}

func (self *GetVantagePoint) CheckConditions() bool {
	return self.context.level != nil &&
		self.context.agent != nil &&
		self.context.closestEnemyPosition != NilCoord
}

func (self *GetVantagePoint) Perform() bool {
	positions := Circle(self.context.closestEnemyPosition, self.context.agent.Range())

	distance := math.MaxFloat64
	self.context.destination = NilCoord
	maps := self.context.level.Map()
	for _, position := range positions {
		if !maps.IsWithinBounds(position) || self.isSightBlocked(maps, position) {
			continue
		}

		distToAgent := Distance(self.context.positionOfAgent, position)
		if distToAgent < distance {
			distance = distToAgent
			self.context.destination = position
		}
	}

	return self.context.destination != NilCoord
}

func (self *GetVantagePoint) isSightBlocked(maps *Map, pos Coord) bool {
	line := Line(self.context.closestEnemyPosition, pos)
	for _, pos := range line {
		if maps.GetCell(pos) == WALL {
			return true
		}
	}

	return false
}
