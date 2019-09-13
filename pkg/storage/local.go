/* LocalStorage can respond Get/Set/Delete and keep persistence by SQLite. */

package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/atpons/m2proxy/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

type LocalStorage struct {
	Store *gorp.DbMap
}

func NewLocalStorage(fileName string) Storage {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTable(Record{}).SetKeys(false, "Key").ColMap("Key").SetUnique(true)
	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}

	if util.Debug > 0 {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	}

	return &LocalStorage{
		Store: dbMap,
	}
}

func (l *LocalStorage) Get(k []byte) (*Record, error) {
	record, err := l.Store.Get(Record{}, string(k))
	if err != nil {
		return nil, ErrStoreInternal
	}
	if r, ok := record.(*Record); !ok {
		return nil, ErrKeyNotFound
	} else {
		return r, nil
	}
}

func (l *LocalStorage) Set(r Record) error {
	record := Record{}
	err := l.Store.SelectOne(&record, "select * from Record where Key=?", r.Key)
	if err != nil {
		return l.Store.Insert(&r)
	}

	if record.CAS != r.CAS || ((record.CAS == 0) && (r.CAS != 0)) {
		return ErrKeyExists
	} else {
		_, err := l.Store.Update(&r)
		return err
	}
}

func (l *LocalStorage) Delete(k []byte) error {
	res, err := l.Store.Delete(&Record{Key: string(k)})

	if res == 0 {
		return ErrKeyNotFound
	}

	if err != nil {
		return ErrStoreInternal
	}

	return nil
}
