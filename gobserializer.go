package main

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct {
	b   bytes.Buffer
	dec *gob.Decoder
}

func newGobSerializer() *gobSerializer {
	s := &gobSerializer{}
	s.dec = gob.NewDecoder(&s.b)
	return s
}

func (g *gobSerializer) Unmarshal(d []byte, o interface{}) error {
	g.b.Reset()
	g.b.Write(d)
	return g.dec.Decode(o)
}