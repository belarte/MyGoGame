package core

// Default values for character's stats.
const (
	DEFAULT_VISIBILITY = 8
	DEFAULT_RANGE      = 1
)

// Character represents a player in the game.
type Character interface {
	Name() string

	MovePoints() float64
	ConsumeMovePoints(float64)
	Reset()

	ActionPoints() int

	Visibility() int
	Range() int
}

// MockCharacter mock
type MockCharacter struct {
	NameMock         string
	MovePointsMock   float64
	ActionPointsMock int
	VisibilityMock   int
	RangeMock        int
}

// Name mock
func (mock *MockCharacter) Name() string {
	return mock.NameMock
}

// MovePoints mock
func (mock *MockCharacter) MovePoints() float64 {
	return mock.MovePointsMock
}

// ConsumeMovePoints mock
func (mock *MockCharacter) ConsumeMovePoints(points float64) {
	mock.MovePointsMock -= points
}

// Reset mock
func (mock *MockCharacter) Reset() {
}

// ActionPoints mock
func (mock *MockCharacter) ActionPoints() int {
	return mock.ActionPointsMock
}

// Visibility mock
func (mock *MockCharacter) Visibility() int {
	return mock.VisibilityMock
}

// Range mock
func (mock *MockCharacter) Range() int {
	return mock.RangeMock
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
