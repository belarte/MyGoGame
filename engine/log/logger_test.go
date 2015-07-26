package log

import "testing"

func TestPrintWithoutChannel(t *testing.T) {
	Print = true
	msg := "This is a test!"
	Log(msg)
}

func TestWithChannel(t *testing.T) {
	Channel = make(chan string)

	msg := "This is a test!"
	go func() { Log(msg) }()

	result := <-Channel
	if msg != result {
		t.Errorf("Expected=%s, got=%s", msg, result)
	}
}
