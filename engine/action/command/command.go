package command

// Command defines the smallest granularity of event that can occur.
// It has two methods, Execute and Revert.
type Command interface {
	Execute()
	Revert()
}

type status struct {
	done bool
}

func newStatus() status {
	return status{done: false}
}

func (c *status) executeIf(command func()) {
	if !c.done {
		command()
		c.done = true
	}
}

func (c *status) revertIf(command func()) {
	if c.done {
		command()
		c.done = false
	}
}

// Fake command.
type Fake struct{}

// Execute fake command.
func (c *Fake) Execute() {}

// Revert fake command.
func (c *Fake) Revert() {}
