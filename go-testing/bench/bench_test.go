package bench_test

import (
	"fmt"
	"go-testing/bench"
	"testing"
)

var blackhole int

func TestFileLen(t *testing.T) {
	result, err := bench.FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("Expected 65204, got", result)
	}
}

func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := bench.FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := bench.FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
