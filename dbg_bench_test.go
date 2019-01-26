package godbg

import (
	"strings"
	"testing"
)

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
