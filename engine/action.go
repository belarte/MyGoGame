package engine

type ActionBaseParameters struct {
	level *Level
	agent *Character
	team  *Team
	logs  [2]string
}

func NewActionBaseParameters(level *Level, agent *Character) ActionBaseParameters {
	team := level.GetTeamOfCharacter(agent)
	return ActionBaseParameters{level, agent, team, [2]string{"", ""}}
}

type Action interface {
	IsDoable() bool
	Perform()
}

type MoveAction struct {
	ActionBaseParameters
	path path
}

func NewMoveAction(lvl *Level, agent *Character, to Coord) *MoveAction {
	finder := NewPathFinder(lvl)
	from := lvl.PositionOfCharacter(agent)
	path := finder.ShortestPath(from, to)
	return &MoveAction{NewActionBaseParameters(lvl, agent), path}
}

func (self *MoveAction) IsDoable() bool {
	if self.path.size() == 0 {
		self.logs[0] = self.agent.Name() + " cannot move: empty path"
		return false
	}

	if self.agent.MovePoints() < self.path.cost() {
		self.logs[0] = self.agent.Name() + " cannot move: not enough move points"
		return false
	}

	self.logs[0] = self.agent.Name() + " can move."
	return true
}

func (self *MoveAction) Perform() {
	for _, step := range self.path.path {
		self.team.MoveCharacter(self.agent, step.coord)
		// TODO: implement cost
		// TODO: implement events
		// TODO: log
	}

	self.logs[1] = self.agent.Name() + " arrived at destination."
}

type AttackAction struct {
	ActionBaseParameters
	target *Character
}

func NewAttackAction(lvl *Level, agent, target *Character) *AttackAction {
	return &AttackAction{NewActionBaseParameters(lvl, agent), target}
}

func (self *AttackAction) IsDoable() bool {
	return false
}

func (self *AttackAction) Perform() {
}
