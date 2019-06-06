package main

import (
	"encoding/json"
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	rawPayload, _ := json.Marshal(&requestLogin{
		Name: "Ben",
		Avatar: "https://server.com/image.jpg",
	})

	for i := 0; i < b.N; i++ {
		var result interface{}
		_ = decode(rawPayload, &result)
	}
}