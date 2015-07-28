package log

import "testing"

func TestPrintWithoutChannel(t *testing.T) {
	doPrint = true
	msg := "This is a test!"
	Log(msg)
}

func TestWithChannel(t *testing.T) {
	channel = make(chan string)

	msg := "This is a test!"
	go func() { Log(msg) }()

	result := <-channel
	if msg != result {
		t.Errorf("Expected=%s, got=%s", msg, result)
	}
}
