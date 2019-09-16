/* LocalStorage can respond Get/Set/Delete and keep persistence by SQLite. */

package storage

import (
	"database/sql"
	"fmt"
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

	if util.Debug > 1 {
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
		fmt.Fprintf(os.Stderr, "local storage: got value: %v\n", r.Value)
		return r, nil
	}
}

func (l *LocalStorage) Set(r Record) (uint64, error) {
	record := Record{}
	err := l.Store.SelectOne(&record, "select * from Record where Key=?", r.Key)

	// If request has CAS, check CAS with selected record
	if (r.CAS > 0) && (record.CAS != r.CAS) {
		return 0, ErrKeyExists
	}

	r.CAS = getCas(r.CAS)

	if err != nil { // No Record and to insert
		return r.CAS, l.Store.Insert(&r)
	}

	_, err = l.Store.Update(&r)
	return r.CAS, err
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

func (l *LocalStorage) Flush() error {
	err := l.Store.TruncateTables()
	if err != nil {
		if util.Debug > 1 {
			fmt.Fprintf(os.Stderr, "localstorage: error truncate tables: %v\n", err)
		}
		return ErrStoreInternal
	}
	return nil
}
