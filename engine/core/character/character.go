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

// Fake Character
type Fake struct {
	FakeName         string
	FakeMovePoints   float64
	FakeConsumeMP    bool
	FakeActionPoints int
	FakeVisibility   int
	FakeRange        int
}

// Name mock
func (mock *Fake) Name() string {
	return mock.FakeName
}

// MovePoints mock
func (mock *Fake) MovePoints() float64 {
	return mock.FakeMovePoints
}

// ConsumeMovePoints mock
func (mock *Fake) ConsumeMovePoints(points float64) bool {
	mock.FakeMovePoints -= points
	return mock.FakeConsumeMP
}

// Reset mock
func (mock *Fake) Reset() {
}

// ActionPoints mock
func (mock *Fake) ActionPoints() int {
	return mock.FakeActionPoints
}

// Visibility mock
func (mock *Fake) Visibility() int {
	return mock.FakeVisibility
}

// Range mock
func (mock *Fake) Range() int {
	return mock.FakeRange
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
