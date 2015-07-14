package core

import (
	. "github.com/belarte/MyGoGame/engine/utils"
)

const (
	MAX_PLAYERS_BY_TEAM = 5
	MAX_PLAYERS         = 2 * MAX_PLAYERS_BY_TEAM
)

type Level struct {
	maps  *Map
	teams []*Team
}

func NewLevel(size Coord, numTeams int) *Level {
	m := NewMap(size)

	teams := make([]*Team, numTeams, numTeams)
	for i := 0; i < numTeams; i++ {
		teams[i] = NewTeam()
	}

	return &Level{m, teams}
}

func (self *Level) Map() *Map {
	return self.maps
}

func (self *Level) AddCharacter(c *Character, pos Coord, team int) bool {
	if !self.maps.IsWithinBounds(pos) {
		return false
	}

	return self.teams[team].AddCharacter(c, pos)
}

func (self *Level) GetTeamOf(char *Character) *Team {
	for _, team := range self.teams {
		if c, _ := team.GetCharacter(char); c != nil {
			return team
		}
	}
	return nil
}

func (self *Level) GetOpponentsOf(char *Character) (result []*Character) {
	team := self.GetTeamOf(char)
	for _, t := range self.teams {
		if t != team {
			result = append(result, t.GetCharacters()...)
		}
	}
	return
}

func (self *Level) PositionOf(c *Character) Coord {
	for _, team := range self.teams {
		if char, coord := team.GetCharacter(c); char != nil {
			return coord
		}
	}

	return Coord{0, 0} // TODO: default return?
}

func (self *Level) IsCharacterAtPosition(pos Coord) bool {
	for _, team := range self.teams {
		if team.IsCharacterAtPosition(pos) {
			return true
		}
	}

	return false
}
