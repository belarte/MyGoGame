package bt

// Sequence execute a list of tasks in the given order.
// It fails when a task fails. It succeeds if all the
// tasks succeed.
type Sequence struct {
	context *context
	tasks   []Task
}

// NewSequence initialise a new Sequence.
func NewSequence(context *context) *Sequence {
	return &Sequence{
		context: context,
		tasks:   make([]Task, 0, 4),
	}
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
func (task *Sequence) Perform() Status {
	for _, t := range task.tasks {
		if !t.CheckConditions() || t.Perform() == failure {
			return failure
		}
	}

	return success
}
