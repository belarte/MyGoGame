package engine

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

/*
func (self *Team) CharacterID(char *Character) (int, error) {
	for i, c := range self.characters {
		if char == c {
			return i, nil
		}
	}

	return -1, errors.New("Character not found")
}

func (self *Team) GetCharacter(id int) *Character {
	return self.characters[id]
}

func (self *Team) GetPosition(id int) Coord {
	return self.positions[id]
}*/

func (self *Team) CharactersCount() int {
	return len(self.characters)
}

func (self *Team) IsCharacterAtPosition(pos Coord) bool {
	for _, p := range self.characters {
		if EqualCoord(p, pos) {
			return true
		}
	}
	return false
}

func (self *Team) IsFull() bool {
	return len(self.characters) == MAX_PLAYERS_BY_TEAM
}
