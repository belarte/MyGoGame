package core

type Character struct {
	name string

	strength     int
	dexterity    int
	agility      int
	intelligence int
	vitality     int

	currentHP, maxHP int
}

func NewCharacter(name string, str, dex, agi, intel, vita int) *Character {
	maxHP := 50 + 10*vita
	return &Character{name, str, dex, agi, intel, vita, maxHP, maxHP}
}

func (self *Character) Name() string {
	return self.name
}

func (self Character) Attack() int {
	return 1 + self.strength
}

func (self Character) MAttack() int {
	return 1 + self.intelligence
}

func (self Character) Defense() int {
	return 1 + self.strength
}

func (self Character) MDefense() int {
	return 1 + self.intelligence
}

func (self Character) MovePoints() int {
	return 1 + self.agility
}

func (self Character) ActionPoints() int {
	return 1 + self.dexterity
}
