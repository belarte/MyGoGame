package ai

// Sequence execute a list of tasks in the given order.
// It fails when a task fails. It succeeds if all the
// tasks succeed.
type Sequence struct {
	context *context
	tasks   []Task
}

// NewSequence initialise a new Sequence.
func NewSequence(context *context) *Sequence {
	return &Sequence{}
}

// Add add a sub task to the list.
func (task *Sequence) Add(t Task) {
	task.tasks = append(task.tasks, t)
}

// CheckConditions verifies that the list of tasks is not empty.
func (task *Sequence) CheckConditions() bool {
	return len(task.tasks) > 0
}

// Perform executes each tasks in the given order.
func (task *Sequence) Perform() bool {
	for _, task := range task.tasks {
		if !task.CheckConditions() || !task.Perform() {
			return false
		}
	}

	return true
}
