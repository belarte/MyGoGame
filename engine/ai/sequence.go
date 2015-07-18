package ai

type Sequence struct {
	context *Context
	tasks   []Task
}

func NewSequence(context *Context) *Sequence {
	return &Sequence{}
}

func (self *Sequence) Add(task Task) {
	self.tasks = append(self.tasks, task)
}

func (self *Sequence) CheckConditions() bool {
	return len(self.tasks) > 0
}

func (self *Sequence) Perform() bool {
	for _, task := range self.tasks {
		if !task.CheckConditions() || !task.Perform() {
			return false
		}
	}

	return true
}
