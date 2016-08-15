package command

// Queue represents a command queue. It is used by Actions to
// queue command that will be processed by the game engine.
type Queue interface {
	Add(c Command)
	ProcessNext()

	Size() int
	Flush()
}

const (
	maxQueueSize = 32
)

// NewQueue returns a new Queue.
func NewQueue() Queue {
	return &queueImpl{
		queue: make(chan Command, 32),
	}
}

type queueImpl struct {
	queue chan Command
}

func (q *queueImpl) Add(c Command) {
	if q.Size() < maxQueueSize {
		q.queue <- c
	}
}

func (q *queueImpl) ProcessNext() {
	if q.Size() > 0 {
		command := <-q.queue
		command.Execute()
	}
}

func (q *queueImpl) Size() int {
	return len(q.queue)
}

func (q *queueImpl) Flush() {
	for q.Size() > 0 {
		<-q.queue
	}
}
