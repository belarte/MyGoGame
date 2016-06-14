package character

// Character represents a player in the game.
type Character interface {
	PositionComponent
	MovePointsComponent
	StatsComponent
}

// Actor represents a actor of the game.
type Actor struct {
	PositionComponent
	MovePointsComponent
	StatsComponent
}
