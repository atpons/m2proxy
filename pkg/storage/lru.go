/* LruStorage can respond Get/Set normally, but cannot respond Delete. */

package storage

import lru "github.com/hashicorp/golang-lru"

type LruStorage struct {
	Store *lru.Cache
}

func NewLruStorage() Storage {
	l, _ := lru.New(128)
	return &LruStorage{
		Store: l,
	}
}

func (l *LruStorage) Get(k []byte) (*Record, error) {
	value, ok := l.Store.Get(string(k))
	if !ok {
		return nil, ErrKeyNotFound
	}
	if v, ok := value.(Record); ok {
		return &v, nil
	}
	return nil, ErrStoreInternal
}

func (l *LruStorage) Set(r Record) (uint64, error) {
	evicted := l.Store.Add(string(r.Key), r)
	if evicted {
		return 0, ErrKeyExists
	}
	return 0, nil
}

func (l *LruStorage) Delete(k []byte) error {
	return nil
}

func (l *LruStorage) Flush() error {
	l.Store.Purge()
	return nil
}
