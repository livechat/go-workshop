package main

import (
	"fmt"
	"bytes"
	"strings"
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
