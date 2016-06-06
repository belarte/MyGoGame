package bt

import (
	"testing"

	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/level"
	"github.com/belarte/MyGoGame/engine/utils"
)

func TestGetVantagePointCheckConditionsNoClosestEnemy(t *testing.T) {
	context := &context{closestEnemyPosition: utils.NilCoord}
	task := NewGetVantagePoint(context)

	if task.CheckConditions() {
		t.Errorf("CheckConditions should return false, context=%+v", context)
	}
}

func TestGetVantagePointCheckConditionsClosestEnemy(t *testing.T) {
	agent := &character.Fake{}
	lvl := level.New(utils.Coord{0, 0}, 0)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{0, 0}}
	task := NewGetVantagePoint(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVantagePointPerformNoObstacle(t *testing.T) {
	agent := &character.Fake{FakeRange: character.DefaultRange}
	lvl := level.New(utils.Coord{1, 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{0, 0}, 0)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{0, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacle(t *testing.T) {
	agent := &character.Fake{FakeRange: character.DefaultRange}
	lvl := level.New(utils.Coord{2, 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{0, 0}, 0)
	lvl.Map().SetCell(utils.Coord{0, 3}, level.WALL)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{1, 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacleAtDistance(t *testing.T) {
	agent := &character.Fake{FakeRange: 3}
	lvl := level.New(utils.Coord{4, 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{0, 0}, 0)
	lvl.Map().SetCell(utils.Coord{0, 3}, level.WALL)
	lvl.Map().SetCell(utils.Coord{1, 3}, level.WALL)
	lvl.Map().SetCell(utils.Coord{2, 3}, level.WALL)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{0, 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{3, 4}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}