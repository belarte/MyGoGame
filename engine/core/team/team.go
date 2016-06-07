package team

import (
	"github.com/belarte/MyGoGame/engine/core/character"
	"github.com/belarte/MyGoGame/engine/utils"
)

// Default max values.
const (
	MaxPlayersByTeam = 5
	MaxPlayers       = 2 * MaxPlayersByTeam
)

// Team represents a team of Character with given positions.
type Team struct {
	characters []character.Character
}

// New returns a new team.
func New() *Team {
	var characters []character.Character
	return &Team{characters}
}

// AddCharacter adds a Character at a given position to the team.
func (team *Team) AddCharacter(c character.Character) bool {
	if team.IsFull() {
		return false
	}

	team.characters = append(team.characters, c)
	return true
}

// Contains checks if the team contains the given character.
func (team *Team) Contains(c character.Character) bool {
	for _, member := range team.characters {
		if member == c {
			return true
		}
	}

	return false
}

// GetCharacters returns a list of all the Characters in the team.
func (team *Team) GetCharacters() (result []character.Character) {
	return team.characters
}

// MoveCharacter moves the given Character to the given position.
//TODO check if actually useful.
func (team *Team) MoveCharacter(char character.Character, pos utils.Coord) {
	for _, c := range team.characters {
		if c == char {
			c.MoveTo(pos)
		}
	}
}

// CharactersCount return the current number of Character in the team.
func (team *Team) CharactersCount() int {
	return len(team.characters)
}

// IsCharacterAtPosition checks if one of the Character is at the given position.
func (team *Team) IsCharacterAtPosition(pos utils.Coord) bool {
	for _, char := range team.characters {
		if char.IsAtPosition(pos) {
			return true
		}
	}
	return false
}

// IsFull checks if the team is at maximum capacity.
func (team *Team) IsFull() bool {
	return len(team.characters) == MaxPlayersByTeam
}
