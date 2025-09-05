package ds

import (
	"strata/api/strata/api"
	"strata/pkg/engine"
	"strconv"
	"time"
)

func Set(e engine.StorageEngine, key, value string) error {
	err := e.Set(key, value)
	if err != nil {
		return err
	}

	EventBus <- &api.StreamResponse{
		Op:        "SET",
		Key:       key,
		Value:     value,
		Timestamp: time.Now().Unix(),
		Version:   1,
	}

	return nil
}

func Get(e engine.StorageEngine, key string) (string, bool, error) {

	ttlStr, ttlFound, _ := e.Get("ttl:" + key)
	if ttlFound {
		ttl, _ := strconv.ParseInt(ttlStr, 10, 64)
		if ttl < time.Now().Unix() {
			e.Del("ttl:" + key)
			e.Del(key)
			return "", false, nil
		}
	}

	return e.Get(key)
}

func Del(e engine.StorageEngine, key string) (bool, error) {
	found, err := e.Del(key)
	return found, err
}

func SetEx(e engine.StorageEngine, key, value string, ttl int64) error {
	if err := e.Set(key, value); err != nil {
		return err
	}

	return e.Set("ttl:"+key, strconv.FormatInt(ttl, 10))
}
