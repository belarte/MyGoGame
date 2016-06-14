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
	lvl := level.New(utils.Coord{X: 0, Y: 0}, 0)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{X: 0, Y: 0}}
	task := NewGetVantagePoint(context)

	if !task.CheckConditions() {
		t.Errorf("CheckConditions should return true, context=%+v", context)
	}
}

func TestGetVantagePointPerformNoObstacle(t *testing.T) {
	agent := &character.Fake{
		StatsComponent:    &character.FakeStatsComponent{FakeRange: character.DefaultRange},
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl := level.New(utils.Coord{X: 1, Y: 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{X: 0, Y: 0}, 0)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{X: 0, Y: 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{X: 0, Y: 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacle(t *testing.T) {
	agent := &character.Fake{
		StatsComponent:    &character.FakeStatsComponent{FakeRange: character.DefaultRange},
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl := level.New(utils.Coord{X: 2, Y: 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{X: 0, Y: 0}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 3}, level.WallCell)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{X: 0, Y: 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{X: 1, Y: 3}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}

func TestGetVantagePointPerformObstacleAtDistance(t *testing.T) {
	agent := &character.Fake{
		StatsComponent:    &character.FakeStatsComponent{FakeRange: 3},
		PositionComponent: &character.FakePositionComponent{},
	}
	lvl := level.New(utils.Coord{X: 4, Y: 5}, 1)
	lvl.AddCharacter(agent, utils.Coord{X: 0, Y: 0}, 0)
	lvl.SetCell(utils.Coord{X: 0, Y: 3}, level.WallCell)
	lvl.SetCell(utils.Coord{X: 1, Y: 3}, level.WallCell)
	lvl.SetCell(utils.Coord{X: 2, Y: 3}, level.WallCell)
	context := &context{agent: agent, lvl: lvl, closestEnemyPosition: utils.Coord{X: 0, Y: 4}}
	task := NewGetVantagePoint(context)

	if !task.Perform() {
		t.Errorf("Perform should return true, context=%+v", context)
	}

	charPosition := context.destination
	expectedPosition := utils.Coord{X: 3, Y: 4}
	if charPosition != expectedPosition {
		t.Errorf("Expected position: %+v, got %+v", expectedPosition, charPosition)
	}
}
