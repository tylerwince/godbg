package godbg

import (
	"bytes"
	"io"
	"os"
	"strings"
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

func BenchmarkWithReverse(b *testing.B) {
	input := "/godbg/cmd/main.go"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := []rune(input)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		result := string(r)
		_ = result
	}
}

func BenchmarkWithLastIndex(b *testing.B) {
	input := "/godbg/cmd/main.go"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := input[strings.LastIndex(input, "/")+1:]
		_ = result
	}
}
