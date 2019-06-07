package main

import (
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	// setup ...
	// tip: use requestLogin or requestBroadcast to demonstrate benchmarking

	for i := 0; i < b.N; i++ {
		// call the decode function
	}
}