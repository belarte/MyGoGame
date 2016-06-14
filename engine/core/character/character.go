package character

// Character represents a player in the game.
type Character interface {
	PositionComponent
	MovePointsComponent
	StatsComponent
}

// Fake Character
type Fake struct {
	PositionComponent
	MovePointsComponent
	StatsComponent
}
