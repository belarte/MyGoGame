package core

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Team represents a team of Character with given positions.
type Team struct {
	characters map[character.Character]utils.Coord
}

// NewTeam returns a new team.
func NewTeam() *Team {
	characters := make(map[character.Character]utils.Coord)
	return &Team{characters}
}

// AddCharacter adds a Character at a given position to the team.
func (team *Team) AddCharacter(c character.Character, pos utils.Coord) bool {
	if team.IsFull() {
		return false
	}

	team.characters[c] = pos

	return true
}

// GetCharacters returns a list of all the Characters in the team.
func (team *Team) GetCharacters() (result []character.Character) {
	for char := range team.characters {
		result = append(result, char)
	}

	return
}

// GetCharacter return a Character and its position.
func (team *Team) GetCharacter(char character.Character) (character.Character, utils.Coord) {
	for c, pos := range team.characters {
		if c == char {
			return c, pos
		}
	}

	return nil, utils.NilCoord
}

// MoveCharacter moves the given Character to the given position.
func (team *Team) MoveCharacter(char character.Character, pos utils.Coord) {
	team.characters[char] = pos
}

// CharactersCount return the current number of Character in the team.
func (team *Team) CharactersCount() int {
	return len(team.characters)
}

// IsCharacterAtPosition checks if one of the Character is at the given position.
func (team *Team) IsCharacterAtPosition(pos utils.Coord) bool {
	for _, p := range team.characters {
		if p == pos {
			return true
		}
	}
	return false
}

// IsFull checks if the team is at maximum capacity.
func (team *Team) IsFull() bool {
	return len(team.characters) == MaxPlayersByTeam
}
