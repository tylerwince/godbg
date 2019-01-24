package godbg

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Dbg will print to either stderr or stdout the output of the expressed passed
func Dbg(exp interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Fprintln(os.Stderr, "Dbg: Unable to parse runtime caller")
		return
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dbg: Unable to open expected file")
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var out string
	i := 1
	for scanner.Scan() {
		if i == line {
			v := scanner.Text()[strings.Index(scanner.Text(), "(")+1 : len(scanner.Text())-strings.Index(reverseString(scanner.Text()), ")")-1]
			out = fmt.Sprintf("[%s:%d] %s = %+v", file[len(file)-strings.Index(reverseString(file), "/"):], line, v, exp)
			break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	switch exp.(type) {
	case error:
		fmt.Fprintln(os.Stderr, out)
	default:
		fmt.Fprintln(os.Stdout, out)
	}
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
