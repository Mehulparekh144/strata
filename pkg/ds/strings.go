package ds

import "strata/pkg/engine"

func Set(e engine.StorageEngine, key, value string) error {
	return e.Set(key, value)
}

func Get(e engine.StorageEngine, key string) (string, bool, error) {
	value, found, err := e.Get(key)
	return value, found, err
}

func Del(e engine.StorageEngine, key string) (bool, error) {
	found, err := e.Del(key)
	return found, err
}
