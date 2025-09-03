package engine

import (
	"github.com/cockroachdb/pebble"
)

type PebbleDBEngine struct {
	db *pebble.DB
}

func NewDBEngine(path string) (*PebbleDBEngine, error) {	
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &PebbleDBEngine{
		db: db,
	}, nil
}

func (e *PebbleDBEngine) Set(key string, value string) error {
	return e.db.Set([]byte(key), []byte(value), pebble.Sync)
}

func (e *PebbleDBEngine) Get(key string) (string, bool, error) {
	val, closer, err := e.db.Get([]byte(key))
	if err != nil {
		if err == pebble.ErrNotFound {
			return "", false, nil
		}
		return "", false, err
	}
	defer closer.Close()
	return string(val), true, nil
}

func (e *PebbleDBEngine) Del(key string) (bool, error) {
	err := e.db.Delete([]byte(key), pebble.Sync)
	if err != nil {
		if err == pebble.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (e *PebbleDBEngine) Close() error {
	e.db.Close()
	return nil
}
