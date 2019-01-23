package godbg

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDbg(t *testing.T) {
	intType := 2
	floatType := 2.1
	strType := "mystring"
	boolType := true

	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	Dbg(intType)
	Dbg(floatType)
	Dbg(strType)
	Dbg(boolType)

	// back to normal state
	w.Close()
	out := <-outC

	want := `[dbg_test.go:25] intType = 2
[dbg_test.go:26] floatType = 2.1
[dbg_test.go:27] strType = mystring
[dbg_test.go:28] boolType = true
`

	if out != want {
		t.Fail()
	}
}
