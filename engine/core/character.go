package core

const (
	DEFAULT_VISIBILITY = 8
	DEFAULT_RANGE      = 1
)

type Character interface {
	Name() string
	MovePoints() int
	ActionPoints() int
	Visibility() int
	Range() int
}

type MockCharacter struct {
	NameMock         string
	MovePointsMock   int
	ActionPointsMock int
	VisibilityMock   int
	RangeMock        int
}

func (self *MockCharacter) Name() string {
	return self.NameMock
}

func (self *MockCharacter) MovePoints() int {
	return self.MovePointsMock
}

func (self *MockCharacter) ActionPoints() int {
	return self.ActionPointsMock
}

func (self *MockCharacter) Visibility() int {
	return self.VisibilityMock
}

func (self *MockCharacter) Range() int {
	return self.RangeMock
}

/*
type Protagonist struct {
	name string

	strength     int
	dexterity    int
	agility      int
	intelligence int
	vitality     int

	currentHP, maxHP int
}

func NewProtagonist(name string, str, dex, agi, intel, vita int) *Protagonist {
	maxHP := 50 + 10*vita
	return &Protagonist{name, str, dex, agi, intel, vita, maxHP, maxHP}
}

func (self *Protagonist) Name() string {
	return self.name
}

func (self Protagonist) Attack() int {
	return 1 + self.strength
}

func (self Protagonist) MAttack() int {
	return 1 + self.intelligence
}

func (self Protagonist) Defense() int {
	return 1 + self.strength
}

func (self Protagonist) MDefense() int {
	return 1 + self.intelligence
}

func (self Protagonist) MovePoints() int {
	return 1 + self.agility
}

func (self Protagonist) ActionPoints() int {
	return 1 + self.dexterity
}

func (self *Protagonist) Visibility() int {
	return DEFAULT_VISIBILITY
}

func (self *Protagonist) Range() int {
	return DEFAULT_RANGE
}
*/
