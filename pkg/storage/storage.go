package storage

import (
	"errors"
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
	CAS   uint32
	Exp   uint32
}

type Storage interface {
	Get([]byte) (*Record, error)
	Set(Record) error
	Delete([]byte) error
}
