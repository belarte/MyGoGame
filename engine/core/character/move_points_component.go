package character

// MovePointsComponent interface
type MovePointsComponent interface {
	MovePoints() float64
	ConsumeMovePoints(float64) bool
	ResetMovePoints()
}

// SimpleMovePointsComponent implements MovePointsComponent
type SimpleMovePointsComponent struct {
	currentMP, baseMP, bonus float64
}

// NewSimpleMovePointsComponent instanciate a simple MovePointsComponent.
func NewSimpleMovePointsComponent(base, bonus float64) *SimpleMovePointsComponent {
	result := &SimpleMovePointsComponent{
		baseMP: base,
		bonus:  bonus,
	}
	result.ResetMovePoints()
	return result
}

// MovePoints returns the current available move points.
func (move *SimpleMovePointsComponent) MovePoints() float64 {
	return move.currentMP
}

// ConsumeMovePoints consumes a given amount of move points.
func (move *SimpleMovePointsComponent) ConsumeMovePoints(points float64) bool {
	if points > move.currentMP {
		return false
	}

	move.currentMP -= points
	return true
}

// ResetMovePoints resets the current move points according to the
// base value and bonus value.
func (move *SimpleMovePointsComponent) ResetMovePoints() {
	move.currentMP = move.baseMP + move.bonus
}

// FakeMovePointsComponent for testing.
type FakeMovePointsComponent struct {
	FakeMovePoints float64
	FakeConsumeMP  bool
}

// MovePoints fake
func (fake *FakeMovePointsComponent) MovePoints() float64 {
	return fake.FakeMovePoints
}

// ConsumeMovePoints fake
func (fake *FakeMovePointsComponent) ConsumeMovePoints(points float64) bool {
	fake.FakeMovePoints -= points
	return fake.FakeConsumeMP
}

// ResetMovePoints fake
func (fake *FakeMovePointsComponent) ResetMovePoints() {
}
