package fib_test

import (
	"testing"

	"github.com/jtprogru/gch/internal/fib"
)

func TestFibRecursive(t *testing.T) {
	testCases := []struct {
		position uint
		gotNum   uint
	}{
		{position: 1, gotNum: 1},
		{position: 2, gotNum: 1},
		{position: 3, gotNum: 2},
		{position: 4, gotNum: 3},
		{position: 5, gotNum: 5},
		{position: 6, gotNum: 8},
		{position: 7, gotNum: 13},
		{position: 8, gotNum: 21},
		{position: 9, gotNum: 34},
		{position: 10, gotNum: 55},
	}

	for _, tt := range testCases {
		t.Run("tt.inputNum", func(t *testing.T) {
			got := fib.Recursive(tt.position)

			if got != tt.gotNum {
				t.Errorf("fib.FibRecursive() = %v, want %v", got, tt.gotNum)
			}

		})
	}

}

func TestFibIterative(t *testing.T) {
	testCases := []struct {
		position uint
		gotNum   uint
	}{
		{position: 1, gotNum: 1},
		{position: 2, gotNum: 1},
		{position: 3, gotNum: 2},
		{position: 4, gotNum: 3},
		{position: 5, gotNum: 5},
		{position: 10, gotNum: 55},
		{position: 17, gotNum: 1597},
		{position: 18, gotNum: 2584},
		{position: 19, gotNum: 4181},
		{position: 20, gotNum: 6765},
		{position: 21, gotNum: 10946},
		{position: 22, gotNum: 17711},

		// 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765
	}

	for _, tt := range testCases {
		t.Run("tt.inputNum", func(t *testing.T) {
			got := fib.Iterative(tt.position)

			if got != tt.gotNum {
				t.Errorf("fib.FibIterative() = %v, want %v", got, tt.gotNum)
			}
		})
	}

}

// Benchmark for Recursive Function.
func BenchmarkFibRecursive(b *testing.B) {
	for range b.N {
		fib.Recursive(uint(20))
	}
}

// Benchmark for Iterative Function.
func BenchmarkFibIterative(b *testing.B) {
	for range b.N {
		fib.Iterative(uint(20))
	}
}
