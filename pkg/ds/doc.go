package ds

import (
	"encoding/json"
	"strata/pkg/engine"

	"github.com/PaesslerAG/jsonpath"
)

func DocSet(e engine.StorageEngine, key, jsonValue string) error {
	var document any
	err := json.Unmarshal([]byte(jsonValue), &document)
	if err != nil {
		return err
	}

	return e.Set("doc:"+key, jsonValue)
}

func DocGet(e engine.StorageEngine, key, path string) (string, bool, error) {
	jsonValue, found, err := e.Get("doc:" + key)
	if err != nil {
		return "", false, err
	}

	if path == "" {
		return jsonValue, found, nil
	}

	var document any
	err = json.Unmarshal([]byte(jsonValue), &document)
	if err != nil {
		return "", false, err
	}

	res, err := jsonpath.Get(path, document)
	if err != nil {
		return "", false, err
	}

	b, _ := json.Marshal(res)
	return string(b), true, nil
}

func DocDel(e engine.StorageEngine, key string) (bool, error) {
	return e.Del("doc:" + key)
}
