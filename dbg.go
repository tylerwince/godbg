package godbg

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Dbg will print to either stderr or stdout the output of the expression passed
func Dbg(exp interface{}) interface{} {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Fprintln(os.Stderr, "Dbg: Unable to parse runtime caller")
		return nil
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dbg: Unable to open expected file")
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var out string
	for i := 1; scanner.Scan(); i++ {
		if i == line {
			v := scanner.Text()[strings.Index(scanner.Text(), "(")+1 : strings.LastIndex(scanner.Text(), ")")]
			out = fmt.Sprintf("[%s:%d] %s = %+v", file[strings.LastIndex(file, "/")+1:], line, v, exp)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return nil
	}
	switch exp.(type) {
	case error:
		fmt.Fprintln(os.Stderr, out)
	default:
		fmt.Fprintln(os.Stdout, out)
	}
	return exp
}
