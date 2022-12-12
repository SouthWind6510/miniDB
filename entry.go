package miniDB

import (
	"time"
)

const headerSize = 17

type OptrType uint8

const (
	PUT OptrType = iota
	DEL
)

type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Type      OptrType
	Timestamp int64
}

func NewEntry(key, value []byte, opType OptrType) *Entry {
	return &Entry{
		Key:       key,
		Value:     value,
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Type:      opType,
		Timestamp: time.Now().Unix(),
	}
}

func (e *Entry) GetSize() uint64 {
	return uint64(headerSize + e.KeySize + e.ValueSize)
}
