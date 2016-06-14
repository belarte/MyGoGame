package level

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/team"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Level represents a level of the game.
// It has a map and a list of teams.
type Level struct {
	Map
	teams []*team.Team
}

// New returns the new level.
func New(size utils.Coord, numTeams int) *Level {
	m := NewMap(size)

	teams := make([]*team.Team, numTeams, numTeams)
	for i := 0; i < numTeams; i++ {
		teams[i] = team.New()
	}

	return &Level{m, teams}
}

// AddActor adds a Actor at a position to the given team.
func (lvl *Level) AddActor(c *character.Actor, pos utils.Coord, team int) bool {
	if !lvl.IsWithinBounds(pos) {
		return false
	}

	c.MoveTo(pos)
	return lvl.teams[team].AddActor(c)
}

// GetTeamOf return the team of the given Actor.
func (lvl *Level) GetTeamOf(char *character.Actor) *team.Team {
	for _, team := range lvl.teams {
		if team.Contains(char) {
			return team
		}
	}
	return nil
}

// GetOpponentsOf returns all the Actors that are not in the team of the given Actor.
func (lvl *Level) GetOpponentsOf(char *character.Actor) (result []*character.Actor) {
	team := lvl.GetTeamOf(char)
	for _, t := range lvl.teams {
		if t != team {
			result = append(result, t.GetActors()...)
		}
	}
	return
}

// IsActorAtPosition checks if the given position is occupied by a Actor.
func (lvl *Level) IsActorAtPosition(pos utils.Coord) bool {
	for _, team := range lvl.teams {
		if team.IsActorAtPosition(pos) {
			return true
		}
	}

	return false
}

// IsObstacleAtPosition checks for obstacles at given position.
func (lvl *Level) IsObstacleAtPosition(pos utils.Coord) bool {
	return lvl.GetCell(pos) == WallCell
}
