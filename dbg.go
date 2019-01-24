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
	switch exp.(type) {
	case error:
		fmt.Fprintln(os.Stderr, out)
	default:
		fmt.Fprintln(os.Stdout, out)
	}
}

func reverseString(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}
