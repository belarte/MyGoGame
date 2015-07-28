package log

import "fmt"

var (
	// Print logs in terminal
	doPrint bool
	// Channel to send logs
	channel chan string
)

func init() {
	doPrint = false
	channel = nil
}

// SetPrint sets if the log should print to console.
func SetPrint(b bool) {
	doPrint = b
}

// Connect connects the logger to an output channel.
func Connect(c chan string) {
	channel = c
}

// Log message on screen or/and send the message through the given channel.
func Log(msg string) {
	if doPrint {
		fmt.Println(">", msg)
	}

	if channel != nil {
		channel <- msg
	}
}
