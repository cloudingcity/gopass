package gopass

import (
	"bytes"
	"sync"
)

var bufPool sync.Pool

func newBuffer() *bytes.Buffer {
	v := bufPool.Get()
	if v == nil {
		return &bytes.Buffer{}
	}
	return v.(*bytes.Buffer)
}

func releaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufPool.Put(b)
}
