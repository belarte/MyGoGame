package log

import "fmt"

var (
	Print   bool
	Channel chan string
)

func init() {
	Print = false
	Channel = nil
}

func Log(msg string) {
	if Print {
		fmt.Println(">", msg)
	}

	if Channel != nil {
		Channel <- msg
	}
}
