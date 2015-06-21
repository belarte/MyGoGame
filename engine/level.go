package engine

import (
	"errors"
)

const (
	MAX_PLAYERS_BY_TEAM = 5
	MAX_PLAYERS         = 2 * MAX_PLAYERS_BY_TEAM
)

type Level struct {
	maps       *Map
	characters []*Character
	positions  []Coord
}

func NewLevel(size Coord) *Level {
	m := NewMap(size)
	characters := make([]*Character, 0, 2*MAX_PLAYERS_BY_TEAM)
	positions := make([]Coord, 0, 2*MAX_PLAYERS_BY_TEAM)

	return &Level{m, characters, positions}
}

func (self *Level) CharactersCount() int {
	return len(self.characters)
}

func (self *Level) AddCharacter(c *Character, pos Coord) error {
	if self.maps.isWithinBounds(pos) {
		return errors.New("Out of bound")
	}

	if len(self.characters) >= MAX_PLAYERS {
		return errors.New("Too much characters already")
	}

	self.characters = append(self.characters, c)
	self.positions = append(self.positions, pos)

	return self.checkConsistency()
}

func (self *Level) CharacterID(char *Character) (int, error) {
	for i, c := range self.characters {
		if char == c {
			return i, nil
		}
	}

	return 0, errors.New("Character not found")
}

func (self *Level) GetCharacter(id int) *Character {
	return self.characters[id]
}

func (self *Level) GetPosition(id int) Coord {
	return self.positions[id]
}

func (self *Level) IsCharacterAtPosition(pos Coord) bool {
	for _, c := range self.positions {
		if EqualCoord(c, pos) {
			return true
		}
	}
	return false
}

func (self *Level) checkConsistency() error {
	if len(self.characters) != len(self.positions) {
		return errors.New("#characters differs from #positions")
	}

	return nil
}
