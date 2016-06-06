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

// Name character
func (character *Fake) Name() string {
	return character.FakeName
}

// MovePoints character
func (character *Fake) MovePoints() float64 {
	return character.FakeMovePoints
}

// ConsumeMovePoints character
func (character *Fake) ConsumeMovePoints(points float64) bool {
	character.FakeMovePoints -= points
	return character.FakeConsumeMP
}

// Reset character
func (character *Fake) Reset() {
}

// Visibility character
func (character *Fake) Visibility() int {
	return character.FakeVisibility
}

// Range character
func (character *Fake) Range() int {
	return character.FakeRange
}
