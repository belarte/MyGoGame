package level

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/core/team"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Level represents a level of the game.
// It has a map and a list of teams.
type Level struct {
	maps  *Map
	teams []*team.Team
}

// NewLevel returns the new level.
func NewLevel(size utils.Coord, numTeams int) *Level {
	m := NewMap(size)

	teams := make([]*team.Team, numTeams, numTeams)
	for i := 0; i < numTeams; i++ {
		teams[i] = team.NewTeam()
	}

	return &Level{m, teams}
}

// Map returns the map.
func (lvl *Level) Map() *Map {
	return lvl.maps
}

// AddCharacter adds a Character at a position to the given team.
func (lvl *Level) AddCharacter(c character.Character, pos utils.Coord, team int) bool {
	if !lvl.maps.IsWithinBounds(pos) {
		return false
	}

	return lvl.teams[team].AddCharacter(c, pos)
}

// GetTeamOf return the team of the given Character.
func (lvl *Level) GetTeamOf(char character.Character) *team.Team {
	for _, team := range lvl.teams {
		if c, _ := team.GetCharacter(char); c != nil {
			return team
		}
	}
	return nil
}

// GetOpponentsOf returns all the Characters that are not in the team of the given Character.
func (lvl *Level) GetOpponentsOf(char character.Character) (result []character.Character) {
	team := lvl.GetTeamOf(char)
	for _, t := range lvl.teams {
		if t != team {
			result = append(result, t.GetCharacters()...)
		}
	}
	return
}

// PositionOf returns the position of the given Character.
func (lvl *Level) PositionOf(c character.Character) utils.Coord {
	for _, team := range lvl.teams {
		if char, coord := team.GetCharacter(c); char != nil {
			return coord
		}
	}

	return utils.NilCoord
}

// IsCharacterAtPosition checks if the given position is occupied by a Character.
func (lvl *Level) IsCharacterAtPosition(pos utils.Coord) bool {
	for _, team := range lvl.teams {
		if team.IsCharacterAtPosition(pos) {
			return true
		}
	}

	return false
}

// IsObstacleAtPosition checks for obstacles at given position.
func (lvl *Level) IsObstacleAtPosition(pos utils.Coord) bool {
	return lvl.maps.GetCell(pos) == WALL
}