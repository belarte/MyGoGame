package log

import "fmt"

var (
	// Print logs in terminal
	Print bool
	// Channel to send logs
	Channel chan string
)

func init() {
	Print = false
	Channel = nil
}

// Log message on screen or/and send the message through the given channel.
func Log(msg string) {
	if Print {
		fmt.Println(">", msg)
	}

	if Channel != nil {
		Channel <- msg
	}
}
