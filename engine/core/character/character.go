package character

// Default values for character's stats.
const (
	DefaultVisibility = 8
	DefaultRange      = 1
	DefaultAP         = 5
)

// Character represents a player in the game.
type Character interface {
	PositionComponent
	MovePointsComponent

	Name() string

	Visibility() int
	Range() int
}

// Fake Character
type Fake struct {
	PositionComponent
	MovePointsComponent

	FakeName       string
	FakeVisibility int
	FakeRange      int
}

// Name character
func (character *Fake) Name() string {
	return character.FakeName
}

// Visibility character
func (character *Fake) Visibility() int {
	return character.FakeVisibility
}

// Range character
func (character *Fake) Range() int {
	return character.FakeRange
}
