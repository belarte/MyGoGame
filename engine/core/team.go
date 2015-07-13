package core

import (
	. "github.com/belarte/MyGoGame/engine/utils"
)

type Team struct {
	characters map[*Character]Coord
}

func NewTeam() *Team {
	characters := make(map[*Character]Coord)
	return &Team{characters}
}

func (self *Team) AddCharacter(c *Character, pos Coord) bool {
	if self.IsFull() {
		return false
	}

	self.characters[c] = pos

	return true
}

func (self *Team) GetCharacters() (result []*Character) {
	for char, _ := range self.characters {
		result = append(result, char)
	}

	return
}

func (self *Team) GetCharacter(char *Character) (*Character, Coord) {
	for c, pos := range self.characters {
		if c == char {
			return c, pos
		}
	}

	return nil, Coord{0, 0}
}

func (self *Team) MoveCharacter(char *Character, pos Coord) {
	self.characters[char] = pos
}

func (self *Team) CharactersCount() int {
	return len(self.characters)
}

func (self *Team) IsCharacterAtPosition(pos Coord) bool {
	for _, p := range self.characters {
		if p == pos {
			return true
		}
	}
	return false
}

func (self *Team) IsFull() bool {
	return len(self.characters) == MAX_PLAYERS_BY_TEAM
}
