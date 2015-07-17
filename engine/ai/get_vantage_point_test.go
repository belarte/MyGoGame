package ai

import (
	. "github.com/belarte/MyGoGame/engine/core"
	. "github.com/belarte/MyGoGame/engine/utils"
	"testing"
)

func TestGetVantagePointCheckConditionsNoClosestEnemy(t *testing.T) {
	context := &Context{closestEnemyPosition: Coord{-1, -1}}
	task := NewGetVantagePoint(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVantagePointCheckConditionsClosestEnemy(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{0, 0}, 0)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 0}}
	task := NewGetVantagePoint(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVantagePointPerformNoObstacle(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{1, 5}, 1)
	level.AddCharacter(agent, Coord{0, 0}, 0)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := Coord{0, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacle(t *testing.T) {
	agent := NewCharacter("", 0, 0, 0, 0, 0)
	level := NewLevel(Coord{2, 5}, 1)
	level.AddCharacter(agent, Coord{0, 0}, 0)
	level.Map().SetCell(Coord{0, 3}, WALL)
	context := &Context{agent: agent, level: level, closestEnemyPosition: Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := Coord{1, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}
