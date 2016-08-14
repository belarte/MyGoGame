package command

// Command defines the smallest granularity of event that can occur.
// It has two methods, Execute and Revert.
type Command interface {
	Execute()
	Revert()
}

type commandStatus struct {
	done bool
}

func newCommandStatus() commandStatus {
	return commandStatus{done: false}
}

func (c *commandStatus) isDone() bool {
	return c.done
}

func (c *commandStatus) markDone() {
	c.done = true
}

func (c *commandStatus) markUndone() {
	c.done = false
}

func (c *commandStatus) executeIf(command func()) {
	if !c.isDone() {
		command()
		c.markDone()
	}
}

func (c *commandStatus) revertIf(command func()) {
	if c.isDone() {
		command()
		c.markUndone()
	}
}
