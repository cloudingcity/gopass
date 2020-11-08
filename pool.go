package gopass

import (
	"bytes"
	"sync"
)

var bufPool sync.Pool

func NewBuffer() *bytes.Buffer {
	v := bufPool.Get()
	if v == nil {
		return &bytes.Buffer{}
	}
	return v.(*bytes.Buffer)
}

func ReleaseBuffer(b *bytes.Buffer) *bytes.Buffer {
	b.Reset()
	bufPool.Put(b)
	return b
}
