/* DebugStorage can respond dummy record. */

package storage

type DebugStorage struct{}

func NewDebugStorage() Storage {
	return &DebugStorage{}
}

func (d *DebugStorage) Get(k []byte) (*Record, error) {
	dummy := Record{
		Key:   string(k),
		Value: []byte{},
		Flag:  0,
		CAS:   0,
		Exp:   0,
	}
	return &dummy, nil
}

func (d *DebugStorage) Set(r Record) (uint64, error) {
	return 0, nil
}

func (d *DebugStorage) Delete(k []byte) error {
	return nil
}

func (d *DebugStorage) Flush() error {
	return nil
}
