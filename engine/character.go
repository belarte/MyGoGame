package engine

type Character struct {
	name string

	strength     uint
	dexterity    uint
	agility      uint
	intelligence uint
	vitality     uint

	currentHP, maxHP  uint
	xp, xpToNextLevel int
	level             int
}

func NewCharacter(name string, str, dex, agi, intel, vita uint) *Character {
	maxHP := 50 + 10*vita
	return &Character{name, str, dex, agi, intel, vita, maxHP, maxHP, 0, 100, 1}
}

func (self *Character) Name() string {
	return self.name
}

func (self Character) Attack() uint {
	return 1 + self.strength
}

func (self Character) MAttack() uint {
	return 1 + self.intelligence
}

func (self Character) Defense() uint {
	return 1 + self.strength
}

func (self Character) MDefense() uint {
	return 1 + self.intelligence
}
