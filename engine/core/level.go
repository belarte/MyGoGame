package core

import (
	"github.com/belarte/MyGoGame/engine/utils"
)

// Default max values.
const (
	MAX_PLAYERS_BY_TEAM = 5
	MAX_PLAYERS         = 2 * MAX_PLAYERS_BY_TEAM
)

// Level represents a level of the game.
// It has a map and a list of teams.
type Level struct {
	maps  *Map
	teams []*Team
}

// NewLevel returns the new level.
func NewLevel(size utils.Coord, numTeams int) *Level {
	m := NewMap(size)

	teams := make([]*Team, numTeams, numTeams)
	for i := 0; i < numTeams; i++ {
		teams[i] = NewTeam()
	}

	return &Level{m, teams}
}

// Map returns the map.
func (lvl *Level) Map() *Map {
	return lvl.maps
}

// AddCharacter adds a Character at a position to the given team.
func (lvl *Level) AddCharacter(c Character, pos utils.Coord, team int) bool {
	if !lvl.maps.IsWithinBounds(pos) {
		return false
	}

	return lvl.teams[team].AddCharacter(c, pos)
}

// GetTeamOf return the team of the given Character.
func (lvl *Level) GetTeamOf(char Character) *Team {
	for _, team := range lvl.teams {
		if c, _ := team.GetCharacter(char); c != nil {
			return team
		}
	}
	return nil
}

// GetOpponentsOf returns all the Characters that are not in the team of the given Character.
func (lvl *Level) GetOpponentsOf(char Character) (result []Character) {
	team := lvl.GetTeamOf(char)
	for _, t := range lvl.teams {
		if t != team {
			result = append(result, t.GetCharacters()...)
		}
	}
	return
}

// PositionOf returns the position of the given Character.
func (lvl *Level) PositionOf(c Character) utils.Coord {
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
