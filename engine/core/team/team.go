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

// Team represents a team of Actor with given positions.
type Team struct {
	characters []*character.Actor
}

// New returns a new team.
func New() *Team {
	var characters []*character.Actor
	return &Team{characters}
}

// AddActor adds a Actor at a given position to the team.
func (team *Team) AddActor(c *character.Actor) bool {
	if team.IsFull() {
		return false
	}

	team.characters = append(team.characters, c)
	return true
}

// Contains checks if the team contains the given character.
func (team *Team) Contains(c *character.Actor) bool {
	for _, member := range team.characters {
		if member == c {
			return true
		}
	}

	return false
}

// GetActors returns a list of all the Actors in the team.
func (team *Team) GetActors() (result []*character.Actor) {
	return team.characters
}

// ActorsCount return the current number of Actor in the team.
func (team *Team) ActorsCount() int {
	return len(team.characters)
}

// IsActorAtPosition checks if one of the Actor is at the given position.
func (team *Team) IsActorAtPosition(pos utils.Coord) bool {
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
