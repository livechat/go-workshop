package main

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(&requestLogin{
		Name: "Ben",
		Avatar: "https://server.com/image.jpg",
	})
	rawPayload := buf.Bytes()

	for i := 0; i < b.N; i++ {
		var result interface{}
		_ = decode(rawPayload, &result)
	}
}