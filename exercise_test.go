package main

import (
	"fmt"
	"bytes"
	"strings"
	"testing"
)

func withPlus(x string) string {
	return x + x + x + x + x + x + x + x + x + x
}

func withSprintf(x string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func withBuffer(x string) string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func withStringBuilder(x string) string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

var sink string

func BenchmarkWithPlus(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withPlus("hello")
	}
	sink = r
}

func BenchmarkWithSprintf(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withSprintf("hello")
	}
	sink = r
}

func BenchmarkWithBuffer(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withBuffer("hello")
	}
	sink = r
}

func BenchmarkWithBuilder(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withStringBuilder("hello")
	}
	sink = r
}