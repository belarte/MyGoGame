package engine

type ActionBaseParameters struct {
	level Level
	agent int
	logs  [2]string
}

func NewActionBaseParameters(level Level, index int) ActionBaseParameters {
	return ActionBaseParameters{level, index, [2]string{"", ""}}
}

type Action interface {
	IsDoable() bool
	Perform()
}

type MoveAction struct {
	ActionBaseParameters
	path []Coord
}

func NewMoveAction(lvl Level, index int, to Coord) *MoveAction {
	path := []Coord{to} // TODO: build path with A*
	return &MoveAction{NewActionBaseParameters(lvl, index), path}
}

func (self *MoveAction) IsDoable() bool {
	// TODO: check action points
	// TODO: check path
	self.logs[0] = self.level.characters[self.agent].name + " can move."
	return true
}

func (self *MoveAction) Perform() {
	for _, position := range self.path {
		self.level.positions[self.agent] = position
		// TODO: implement events
	}

	self.logs[1] = self.level.characters[self.agent].name + " arrived at destination."
}

type AttackAction struct {
	ActionBaseParameters
	indexOfTarget int
}

func NewAttackAction(lvl Level, index int, target int) *AttackAction {
	return &AttackAction{NewActionBaseParameters(lvl, index), target}
}

func (self *AttackAction) IsDoable() bool {
	return false
}

func (self *AttackAction) Perform() {
}
