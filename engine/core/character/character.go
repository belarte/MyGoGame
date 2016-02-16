package character

// Default values for character's stats.
const (
	DefaultVisibility = 8
	DefaultRange      = 1
	DefaultAP         = 5
)

// Character represents a player in the game.
type Character interface {
	Name() string

	MovePoints() float64
	ConsumeMovePoints(float64) bool
	Reset()

	ActionPoints() int

	Visibility() int
	Range() int
}

// Protagonist is the concrete implementation of Character.
type Protagonist struct {
	movePointsHandler

	name string
}

// New returns a new Protagonist.
func New() *Protagonist {
	return &Protagonist{movePointsHandler{1, 2, 3}, "Bob"}
}

// Name returns the name of the character.
func (p *Protagonist) Name() string {
	return p.name
}

// ActionPoints returns the current action points.
func (p *Protagonist) ActionPoints() int {
	return DefaultAP
}

// Visibility returns the visibility.
func (p *Protagonist) Visibility() int {
	return DefaultVisibility
}

// Range returns the range.
func (p *Protagonist) Range() int {
	return DefaultRange
}

type movePointsHandler struct {
	currentMP, baseMP, bonus float64
}

func newMovePointsHandler(base, bonus float64) movePointsHandler {
	result := movePointsHandler{
		baseMP: base,
		bonus:  bonus,
	}
	result.Reset()
	return result
}

func (handler *movePointsHandler) MovePoints() float64 {
	return handler.currentMP
}

func (handler *movePointsHandler) ConsumeMovePoints(points float64) bool {
	if points > handler.currentMP {
		return false
	}

	handler.currentMP -= points
	return true
}

func (handler *movePointsHandler) Reset() {
	handler.currentMP = handler.baseMP + handler.bonus
}

// Mock Character
type Mock struct {
	NameMock         string
	MovePointsMock   float64
	ConsumeMPMock    bool
	ActionPointsMock int
	VisibilityMock   int
	RangeMock        int
}

// Name mock
func (mock *Mock) Name() string {
	return mock.NameMock
}

// MovePoints mock
func (mock *Mock) MovePoints() float64 {
	return mock.MovePointsMock
}

// ConsumeMovePoints mock
func (mock *Mock) ConsumeMovePoints(points float64) bool {
	mock.MovePointsMock -= points
	return mock.ConsumeMPMock
}

// Reset mock
func (mock *Mock) Reset() {
}

// ActionPoints mock
func (mock *Mock) ActionPoints() int {
	return mock.ActionPointsMock
}

// Visibility mock
func (mock *Mock) Visibility() int {
	return mock.VisibilityMock
}

// Range mock
func (mock *Mock) Range() int {
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
*/
