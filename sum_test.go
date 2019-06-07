package main

import "testing"

func sum(i, j int) int {
	return i + j
}

var sink int

func BenchmarkSum(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = sum(i, i) // prevent eliminating the function call
	}
	sink = r // prevent eliminating the benchmark itself
}