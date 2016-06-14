package character

// Default values for character's stats.
const (
	DefaultVisibility = 8
	DefaultRange      = 1
	DefaultAP         = 5
)

// StatsComponent represents all basic stats of a character.
type StatsComponent interface {
	Name() string
	Visibility() int
	Range() int
}

// FakeStatsComponent represents a fake StatsComponent.
type FakeStatsComponent struct {
	FakeName       string
	FakeVisibility int
	FakeRange      int
}

// Name component
func (c *FakeStatsComponent) Name() string {
	return c.FakeName
}

// Visibility component
func (c *FakeStatsComponent) Visibility() int {
	return c.FakeVisibility
}

// Range component
func (c *FakeStatsComponent) Range() int {
	return c.FakeRange
}
