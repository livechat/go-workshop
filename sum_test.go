package main

import "testing"

func sum(i, j int) int {
	return i + j
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(i, i)
	}
}