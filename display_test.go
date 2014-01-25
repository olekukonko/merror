package merror

import (
	"bytes"
	"testing"
)

func TestTab(t *testing.T) {
	var err Errors
	err.Append(errorList...)
	buf := new(bytes.Buffer)
	display := err.Display(buf)
	display.Tab()
	if buf.String() == "" {
		t.Error("Tab Result Should not be empty")
	}
}


func TestLine(t *testing.T) {
	var err Errors
	err.Append(errorList...)
	buf := new(bytes.Buffer)
	display := err.Display(buf)
	display.Line()
	if buf.String() == "" {
		t.Error("Line Result Shwould not be empty")
	}
}
