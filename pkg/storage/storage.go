package storage

import (
	"errors"
	"math/rand"
	"time"
)

var (
	ErrKeyNotFound   = errors.New("Key Not Found")
	ErrKeyExists     = errors.New("Key Exists")
	ErrStoreInternal = errors.New("Internal Error")
)

type Record struct {
	Key   string
	Value []byte
	Flag  uint32
	CAS   uint64
	Exp   uint32
}

type Storage interface {
	Get([]byte) (*Record, error)
	Set(Record) (uint64, error)
	Delete([]byte) error
	Flush() error
}

func getCas(prev uint64) uint64 {
	for {
		rand.Seed(time.Now().UnixNano())
		num := uint64(rand.Int63())
		if prev != num {
			return num
		}
	}
}
